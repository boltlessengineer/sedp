<p align="center">
    <img src="https://i.imgur.com/jAeH66B.png">
</p>

![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
![C++](https://img.shields.io/badge/c++-%2300599C.svg?style=for-the-badge&logo=c%2B%2B&logoColor=white)

![Arduino](https://img.shields.io/badge/-Arduino-00979D?style=for-the-badge&logo=Arduino&logoColor=white)
![Raspberry Pi](https://img.shields.io/badge/-RaspberryPi-C51A4A?style=for-the-badge&logo=Raspberry-Pi)

![Android Studio](https://img.shields.io/badge/Android_Studio-3DDC84?style=for-the-badge&logo=AndroidStudio&logoColor=white)
![Java](https://img.shields.io/badge/Java-ED8B00?style=for-the-badge&logo=java&logoColor=white)

# S.E.D.P.

> **S** moke\
> **E** mission\
> **D** etecting\
> **P** rogram

> ~~쌤 이새끼 담배 펴요~~

## 개요

SEDP는 오픈소스 담배 연기 감지 시스템입니다.

전국의 어떤 공학 동아리에서든지 쉽게 따라 만들고 설치하여 교내 흡연 문제를 종식시키는 것을 목적으로 하고 있습니다.

## 개발 가이드

1. Install go 1.15 or later.
2. Install ngrok.
3. Move to `/broker/goMQTTBroker` folder and run `./main --host localhost -p 1883`
4. Move to `/server` folder and run `go run ./server.go`
5. Run `./ngrok http 8000`
6. Open `<ngrok host url>/static/` to view the frontend.

더 자세한 가이드 : [SEDP Wiki](https://github.com/boltlessengineer/bopyung-sedp/wiki)

## 문의

개발자 이메일 : [boltlessengineer@gmail.com](mailto://boltlessengineer@gmail.com)
