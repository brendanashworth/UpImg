# UpImg
> UpImg is an open source image-upload tool, based in the language of Go. Images are stored in a MySQL database and encoded/decoded via Base64.

### Compilation
```bash
go get github.com/boboman13/UpImg
cd $GOPATH/UpImg/src/main

go build
./main
```

### Contributing
Simply fork the project, edit changes to your local repository, and submit a pull request to contribute to the project. We follow basic coding standards, and keep in mind the goals of the project before submitting a pull request.

If the fork, add-on, or modification to the code does not reflect on the project's goals, the pull request will (probably) not be accepted. If this happens, in its stead, you may notify the UpImg manager (@boboman13) to have the fork noted and linked to in the readme. Note: follow the LICENSE!

### Project Core Goals
* **Cleanliness**: Keep the source code clean and easy to read. In addition, keep the user interface very easy to manipulate and use. Try to make image uploads as seamless as possible.
* **Speed**: Make sure UpImg is fast. It'll be bottlenecked by network speeds, but lets try and cut out everywhere we can with speed. (This includes Javascript UI)
* **Usability**: There should be no learning curve for the software. The config should be kept minimal and the upload should be a one-stop shop.