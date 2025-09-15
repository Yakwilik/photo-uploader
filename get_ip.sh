#!/bin/bash

echo "🔍 Получение IP-адреса для подключения к серверу..."
echo ""

# Получаем IP-адрес
IP=$(ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | head -1)

if [ -z "$IP" ]; then
    echo "❌ Не удалось определить IP-адрес"
    echo "💡 Убедитесь, что вы подключены к Wi-Fi сети"
    exit 1
fi

echo "🌐 Ваш IP-адрес: $IP"
echo ""
echo "📱 Для подключения с мобильных устройств:"
echo "   http://$IP:8080"
echo ""
echo "🧪 Тестовая страница:"
echo "   http://$IP:8080/test"
echo ""
echo "💡 Убедитесь, что устройство подключено к той же Wi-Fi сети"
echo "✅ Совместимо с iOS 9.3.5+, Android 4.0+, старыми браузерами"
