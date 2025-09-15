package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/khasbulatabdullin/photo-uploader/internal/version"
)

const (
	uploadDir   = "./uploads"
	maxFileSize = 50 << 20 // 50 MB для множественных файлов
)

// Структура для передачи данных в шаблон
type PageData struct {
	Message string
	Error   string
}

// Функция для получения локального IP-адреса
func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Printf("Ошибка получения IP-адреса: %v", err)
		return "localhost"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// ensureUploadDir создает директорию для загрузок с проверками
func ensureUploadDir() error {
	// Проверяем, существует ли директория
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		log.Printf("Директория %s не существует, создаем...", uploadDir)

		// Создаем директорию с правами 0755 (rwxr-xr-x)
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return fmt.Errorf("не удалось создать директорию %s: %v", uploadDir, err)
		}

		log.Printf("✅ Директория %s успешно создана", uploadDir)
	} else if err != nil {
		// Другая ошибка при проверке директории
		return fmt.Errorf("ошибка при проверке директории %s: %v", uploadDir, err)
	} else {
		// Директория существует, проверяем права доступа
		if err := checkDirPermissions(uploadDir); err != nil {
			return fmt.Errorf("проблемы с правами доступа к директории %s: %v", uploadDir, err)
		}
		log.Printf("✅ Директория %s готова к использованию", uploadDir)
	}

	return nil
}

// checkDirPermissions проверяет права доступа к директории
func checkDirPermissions(dir string) error {
	// Проверяем, можем ли мы читать директорию
	if _, err := os.ReadDir(dir); err != nil {
		return fmt.Errorf("нет прав на чтение директории: %v", err)
	}

	// Проверяем, можем ли мы создавать файлы в директории
	testFile := filepath.Join(dir, ".test_write_permission")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return fmt.Errorf("нет прав на запись в директорию: %v", err)
	}

	// Удаляем тестовый файл
	os.Remove(testFile)

	return nil
}

func main() {
	// Обработка флагов командной строки
	var showVersion = flag.Bool("version", false, "Show version information")
	var port = flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	if *showVersion {
		version.Print()
		return
	}

	// Создаем директорию для загрузок, если она не существует
	if err := ensureUploadDir(); err != nil {
		log.Fatal("Не удалось создать директорию для загрузок:", err)
	}

	// Настраиваем маршруты
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/uploads/", serveFileHandler)
	http.HandleFunc("/test", testHandler)

	// Статические файлы
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Получаем IP-адрес
	ip := getLocalIP()

	fmt.Println("=" + strings.Repeat("=", 60))
	fmt.Println("🚀 СЕРВЕР ЗАПУЩЕН")
	fmt.Println("=" + strings.Repeat("=", 60))
	fmt.Printf("📍 Локальный адрес: http://localhost:%s\n", *port)
	fmt.Printf("🌐 Сетевой адрес:   http://%s:%s\n", ip, *port)
	fmt.Println("=" + strings.Repeat("=", 60))
	fmt.Println("📱 Для подключения с мобильных устройств:")
	fmt.Printf("   Откройте браузер и введите: http://%s:%s\n", ip, *port)
	fmt.Println("=" + strings.Repeat("=", 60))
	fmt.Println("✅ Совместим со старыми браузерами (IE9+, Safari 5+, Chrome 20+)")
	fmt.Println("📱 Поддерживает iOS 9.3.5, Android 4.0+, старые планшеты")
	fmt.Println("📁 Поддерживает загрузку любых файлов")
	fmt.Println("🔄 Множественный выбор файлов")
	fmt.Println("=" + strings.Repeat("=", 60))

	log.Fatal(http.ListenAndServe("0.0.0.0:"+*port, nil))
}

// Обработчик тестовой страницы
func testHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test_ios.html")
}

// Обработчик главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="format-detection" content="telephone=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
    <title>Загрузка файлов</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            text-align: center;
            margin-bottom: 30px;
        }
        .upload-form {
            text-align: center;
        }
        .file-input-wrapper {
            position: relative;
            display: inline-block;
            margin: 20px 0;
        }
        .file-input {
            position: absolute;
            left: -9999px;
        }
        .file-input-label {
            display: inline-block;
            padding: 12px 24px;
            background-color: #007AFF;
            color: white;
            border-radius: 8px;
            cursor: pointer;
            font-size: 16px;
            transition: background-color 0.3s;
        }
        .file-input-label:hover {
            background-color: #0056CC;
        }
        .file-input-label:active {
            background-color: #004499;
        }
        .submit-btn {
            background-color: #34C759;
            color: white;
            border: none;
            padding: 12px 24px;
            border-radius: 8px;
            font-size: 16px;
            cursor: pointer;
            margin-top: 20px;
            transition: background-color 0.3s;
        }
        .submit-btn:hover {
            background-color: #28A745;
        }
        .submit-btn:active {
            background-color: #1E7E34;
        }
        .submit-btn:disabled {
            background-color: #ccc;
            cursor: not-allowed;
        }
        .message {
            margin: 20px 0;
            padding: 15px;
            border-radius: 8px;
            text-align: center;
        }
        .success {
            background-color: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        .error {
            background-color: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        .file-info {
            margin: 10px 0;
            font-size: 14px;
            color: #666;
        }
        .progress {
            width: 100%;
            height: 20px;
            background-color: #f0f0f0;
            border-radius: 10px;
            overflow: hidden;
            margin: 20px 0;
            display: none;
        }
        .progress-bar {
            height: 100%;
            background-color: #007AFF;
            width: 0%;
            transition: width 0.3s;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>📁 Загрузка файлов</h1>
        <p style="text-align: center; color: #666; margin-bottom: 20px;">Выберите один или несколько файлов для загрузки</p>
        
        {{if .Message}}
        <div class="message success">{{.Message}}</div>
        {{end}}
        
        {{if .Error}}
        <div class="message error">{{.Error}}</div>
        {{end}}
        
        <form class="upload-form" action="/upload" method="post" enctype="multipart/form-data" id="uploadForm">
            <div class="file-input-wrapper">
                <input type="file" name="photo" id="photos" class="file-input" multiple required>
                <label for="photos" class="file-input-label">Выбрать файлы</label>
            </div>
            
            <div class="file-info" id="fileInfo" style="display: none;">
                <span id="fileName"></span> (<span id="fileSize"></span>)
            </div>
            
            <div class="progress" id="progress">
                <div class="progress-bar" id="progressBar"></div>
            </div>
            <div id="progressText" style="text-align: center; margin: 10px 0; font-size: 14px; color: #666;"></div>
            
            <button type="submit" class="submit-btn" id="submitBtn">Загрузить</button>
        </form>
    </div>

    <script>
        // Совместимость с iOS 9.3.5
        document.getElementById('photos').addEventListener('change', function(e) {
            var files = e.target.files;
            var fileInfo = document.getElementById('fileInfo');
            var fileName = document.getElementById('fileName');
            var fileSize = document.getElementById('fileSize');
            
            if (files && files.length > 0) {
                if (files.length === 1) {
                    fileName.textContent = files[0].name;
                    fileSize.textContent = formatFileSize(files[0].size);
                } else {
                    var totalSize = 0;
                    for (var i = 0; i < files.length; i++) {
                        totalSize += files[i].size;
                    }
                    fileName.textContent = files.length + ' файлов выбрано';
                    fileSize.textContent = formatFileSize(totalSize);
                }
                fileInfo.style.display = 'block';
            } else {
                fileInfo.style.display = 'none';
            }
        });
        
        function formatFileSize(bytes) {
            if (bytes === 0) return '0 Bytes';
            var k = 1024;
            var sizes = ['Bytes', 'KB', 'MB', 'GB'];
            var i = Math.floor(Math.log(bytes) / Math.log(k));
            return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
        }
        
        // Обработка отправки формы
        document.getElementById('uploadForm').addEventListener('submit', function(e) {
            e.preventDefault(); // Предотвращаем стандартную отправку формы
            
            var files = document.getElementById('photos').files;
            if (files.length === 0) {
                alert('Выберите файлы для загрузки');
                return;
            }
            
            var submitBtn = document.getElementById('submitBtn');
            var progress = document.getElementById('progress');
            var progressBar = document.getElementById('progressBar');
            var progressText = document.getElementById('progressText');
            
            submitBtn.disabled = true;
            submitBtn.textContent = 'Загрузка...';
            progress.style.display = 'block';
            progressText.textContent = 'Загрузка файлов...';
            
            var uploadedCount = 0;
            var totalFiles = files.length;
            var errors = [];
            
            // Функция для загрузки одного файла
            function uploadFile(file, index) {
                var formData = new FormData();
                formData.append('photo', file);
                
                var xhr = new XMLHttpRequest();
                xhr.open('POST', '/upload', true);
                
                xhr.onload = function() {
                    uploadedCount++;
                    var progressPercent = (uploadedCount / totalFiles) * 100;
                    progressBar.style.width = progressPercent + '%';
                    progressText.textContent = 'Загружено ' + uploadedCount + ' из ' + totalFiles + ' файлов';
                    
                    if (xhr.status === 200) {
                        // Успешная загрузка
                        console.log('Файл ' + file.name + ' загружен успешно');
                    } else {
                        // Ошибка загрузки
                        errors.push('Ошибка загрузки файла: ' + file.name);
                    }
                    
                    // Если все файлы загружены
                    if (uploadedCount === totalFiles) {
                        submitBtn.disabled = false;
                        submitBtn.textContent = 'Загрузить';
                        progressText.textContent = 'Загрузка завершена';
                        
                        if (errors.length === 0) {
                            alert('Все файлы загружены успешно!');
                            window.location.reload();
                        } else {
                            alert('Загрузка завершена с ошибками:\\n' + errors.join('\\n'));
                        }
                    }
                };
                
                xhr.onerror = function() {
                    uploadedCount++;
                    errors.push('Ошибка сети при загрузке файла: ' + file.name);
                    
                    if (uploadedCount === totalFiles) {
                        submitBtn.disabled = false;
                        submitBtn.textContent = 'Загрузить';
                        alert('Загрузка завершена с ошибками:\\n' + errors.join('\\n'));
                    }
                };
                
                xhr.send(formData);
            }
            
            // Загружаем файлы по одному
            for (var i = 0; i < files.length; i++) {
                uploadFile(files[i], i);
            }
        });
    </script>
</body>
</html>`

	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}

	data := PageData{}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Ошибка выполнения шаблона", http.StatusInternalServerError)
		return
	}
}

// Обработчик загрузки файлов
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Ограничиваем размер запроса
	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)

	// Парсим multipart form
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		showErrorPage(w, "Файлы слишком большие. Максимальный размер: 50MB")
		return
	}

	// Получаем файл
	file, handler, err := r.FormFile("photo")
	if err != nil {
		showErrorPage(w, "Ошибка получения файла: "+err.Error())
		return
	}
	defer file.Close()

	// Проверяем, что файл не пустой
	if handler.Size == 0 {
		showErrorPage(w, "Выбранный файл пустой")
		return
	}

	// Убеждаемся, что директория uploads существует и доступна
	if err := ensureUploadDir(); err != nil {
		showErrorPage(w, "Ошибка доступа к директории загрузок: "+err.Error())
		return
	}

	// Создаем уникальное имя файла
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filepath := filepath.Join(uploadDir, filename)

	// Создаем файл на диске
	dst, err := os.Create(filepath)
	if err != nil {
		showErrorPage(w, "Ошибка создания файла: "+err.Error())
		return
	}
	defer dst.Close()

	// Копируем содержимое файла
	_, err = io.Copy(dst, file)
	if err != nil {
		showErrorPage(w, "Ошибка сохранения файла: "+err.Error())
		return
	}

	// Показываем страницу успеха
	showSuccessPage(w, fmt.Sprintf("Файл '%s' успешно загружен!", handler.Filename))
}

// Обработчик для отдачи статических файлов
func serveFileHandler(w http.ResponseWriter, r *http.Request) {
	// Безопасность: проверяем, что путь не содержит ".."
	if strings.Contains(r.URL.Path, "..") {
		http.Error(w, "Запрещенный путь", http.StatusForbidden)
		return
	}

	// Убеждаемся, что директория uploads существует
	if err := ensureUploadDir(); err != nil {
		http.Error(w, "Ошибка доступа к директории загрузок", http.StatusInternalServerError)
		return
	}

	filepath := filepath.Join(".", r.URL.Path)
	http.ServeFile(w, r, filepath)
}

// Показывает страницу с ошибкой
func showErrorPage(w http.ResponseWriter, errorMsg string) {
	tmpl := `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ошибка загрузки</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            text-align: center;
        }
        .error {
            background-color: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
            padding: 15px;
            border-radius: 8px;
            margin: 20px 0;
        }
        .back-btn {
            background-color: #007AFF;
            color: white;
            border: none;
            padding: 12px 24px;
            border-radius: 8px;
            font-size: 16px;
            cursor: pointer;
            text-decoration: none;
            display: inline-block;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>❌ Ошибка</h1>
        <div class="error">{{.Error}}</div>
        <a href="/" class="back-btn">← Вернуться</a>
    </div>
</body>
</html>`

	t, _ := template.New("error").Parse(tmpl)
	data := PageData{Error: errorMsg}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
	}
}

// Показывает страницу успеха
func showSuccessPage(w http.ResponseWriter, successMsg string) {
	tmpl := `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Успешная загрузка</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            text-align: center;
        }
        .success {
            background-color: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
            padding: 15px;
            border-radius: 8px;
            margin: 20px 0;
        }
        .back-btn {
            background-color: #34C759;
            color: white;
            border: none;
            padding: 12px 24px;
            border-radius: 8px;
            font-size: 16px;
            cursor: pointer;
            text-decoration: none;
            display: inline-block;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>✅ Успешно!</h1>
        <div class="success">{{.Message}}</div>
        <a href="/" class="back-btn">← Загрузить еще</a>
    </div>
</body>
</html>`

	t, _ := template.New("success").Parse(tmpl)
	data := PageData{Message: successMsg}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
	}
}
