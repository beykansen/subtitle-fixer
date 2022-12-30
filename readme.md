# Toplu Türkçe Altyazı Karakter Sorunu Düzeltici
Bilgisayarınıza ```go``` kurduktan sonra, aşağıdaki komutu çalıştırarak ```subtitle-fixer``` programını yükleyebilirsiniz.
```
go install github.com/beykansen/subtitle-fixer@latest
```
Terminalden programı direkt olarak çalıştırabilmek için ```GOPATH``` environment variable olarak setlenmiş olmalı. https://github.com/golang/go/wiki/SettingGOPATH#setting-gopath

Programı yükledikten sonra ```path``` kısmını kendinize göre özelleştirerek aşağıdaki gibi çalıştırın. Eğer ```path``` vermezseniz, bulunduğunuz klasörden çalışmaya başlayacaktır.
```
subtitle-fixer --path "C:\Filmler"
```