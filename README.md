# GotX

Gotx的中文文档可以从[这里](https://www.jianshu.com/p/d429a3aa0b1c)看起。

A Golang interpreter based on [Yaegi](https://github.com/containous/yaegi).

Download the latest release binaries [here](http://topget.org/gox.html) or [here](https://github.com/topxeq/gotx/releases).

Directly run Golang-like code files. Almost the same syntax as Golang. Installation of Golang is not required to run GotX code, it's convenient especially for running/debugging on remote servers.

Most of the standard libraries and some useful 3rd-party libraries are included.

Support GUI programming with OpenGL or Sciter library.

A command-line only version (GotXc) is also available.

Currently supported 3rd-party libraies: github.com/topxeq/tk(many useful functions), github.com/topxeq/sqltk(SQL databases), github.com/beevik/etree(XML processing), github.com/AllenDang/giu(OpenGL GUI), github.com/360EntSecGroup-Skylar/excelize(Excel file handling), github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi(Send Aliyun SMS), github.com/boombuler/barcode & github.com/boombuler/barcode/qr(generate QR Code), github.com/sciter-sdk/go-sciter & github.com/sciter-sdk/go-sciter/window(support cross-platform HTML/CSS/JavaScript GUI programming), github.com/topxeq/xiaoxian(NLP related toolkit), github.com/topxeq/doc2vec/doc2vec(document to vector), github.com/stretchr/objx(another JSON toolkit), github.com/topxeq/blink(embedable Web Browser from miniblink), .

If more packages are required, please compile the source code, add the packages you need, refer to the examples of the packages already supported.

### Install

Download it from [here](http://topget.org/gox.html). Extract the download zip file, and put the executable anywhere but best in the system PATH. Run it from a command-line console.

And install the Sciter library correspond to your OS before using Sciter to GUI programming with Gotx. It's simple, actually only one dll file is needed for running the GUI apps or distributed it. Refer to the installation documents [here](https://github.com/sciter-sdk/go-sciter), only the first two steps are neccessary. Or, you can run command 'gotx -initgui' to automatically download the latest sciter.dll file to application path of Gotx.
