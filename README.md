[![Contributors][contributors-shield]][contributors-url]
  [![Issues][issues-shield]][issues-url]
  [![MIT License][license-shield]][license-url]

<p align="center">
  <h3 align="center">Temp File Transfer Web Application</h3>
  <br/>
  


  <p align="center">
    Simple temporary file upload and transfer web application coding with Go language.
    <br />
    <a href="https://go.dev/"><strong>Explore the Golang »</strong></a>
    <br />
    <br />
    <a href="https://alya-temp-file.herokuapp.com/"><strong>Live Demo</strong></a>
    <br />
    <br />
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#todo">TODO</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

## About The Project
This project aiming to create a simple temporary file trasnfer app for general purposes. With this app you could upload 
file to service then retrieve from unique key for specific (1 minute) time validity. Demo app is listening on (https://alya-temp-file.herokuapp.com)
. 

### Built With

- [Golang](https://go.dev/) for coding
- [Docker](https://www.docker.com) for containerizing.
- [Heroku](https://heroku.com/) for serving application on the net

## Getting Started
### Prerequisites

- Golang
  ```
  Install latest Golang  
  https://go.dev 
  ```  
- Docker
  ```sh
  Install docker on your OS  
  https://docs.docker.com/get-docker/  
  ```
- Heroku
  ```
  Quick look website  
  https://www.heroku.com 
  ```  


### Installation

1. Clone the repo:
   ```sh
   git clone https://github.com/AlperRehaYAZGAN/temp-file-transfer-app.git  
   cd temp-file-transfer-app
   ```
1. Run app directly:
   ```sh
   go build -o ./bin/myexeapp
   ./bin/myexeapp
   ```
2. Build docker image:
   ```sh
   docker build -t alperreha/tempfiletransfer:1.0.0
   ```
3. Run Docker container:
   ```sh
   docker run --name alya-temp-file -p 9090:9090 -d alperreha/tempfiletransfer:1.0.0
   ```

## Usage

This simple app has a two endpoint to handle whole process. If we assert server is listening on 9090 port, example requests are:

- GET / : HTML form for upload file  
- GET /get/:file-id : Returns file by given file-id   
- POST /upload-one : Form Data myfile for uploading file and returns file key to access  ,

## Roadmap

See the [open issues](https://github.com/AlperRehaYAZGAN/temp-file-transfer-app/issues) for a list of proposed features (and known issues).

## TODO  
- [X] JWT encode,decode and verify  
- [-] Custom TCP Transport for microservices  
- [-] 

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Alper Reha YAZGAN - [@alperreha](https://twitter.com/alperreha) - alper@yazgan.xyz

Project Link: [https://github.com/AlperRehaYAZGAN/temp-file-transfer-app](https://github.com/AlperRehaYAZGAN/temp-file-transfer-app)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/AlperRehaYAZGAN/temp-file-transfer-app.svg?style=for-the-badge
[contributors-url]: https://github.com/AlperRehaYAZGAN/temp-file-transfer-app/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/AlperRehaYAZGAN/temp-file-transfer-app.svg?style=for-the-badge
[issues-url]: https://github.com/AlperRehaYAZGAN/temp-file-transfer-app/issues
[license-shield]: https://img.shields.io/github/license/AlperRehaYAZGAN/temp-file-transfer-app.svg?style=for-the-badge
[license-url]: https://github.com/AlperRehaYAZGAN/temp-file-transfer-app/blob/master/LICENSE.txt