# social_networking_app_golang

<a id="readme-top"></a>

[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



The main objective of this project is to demonstrate my skills in backend development in the Go language.
For this I developed a social network that was divided into two parts: API and Web Application.
The API is the part responsible for making direct interactions with the database and is organized as follows starting from the root `api/`(with a brief explanation of the function of each package):
```
.
├── go.mod
├── go.sum
├── main.go                                     (Main file)
├── .env 
├── sql
│   └── sql.sql                                 (SQL script to create the database)
└── src
    ├── autenticacao                            (Package with functions that manipulate the JWT token)
    ├── banco                                   (Package that only opens the connection to the bank)
    ├── config                                  (Loads environment variables)
    ├── controllers                             (Package responsible for handling API requests)
    ├── middlewares                             (Middleware for logging of request and authentication)
    ├── modelos                                 (Defines the data structures of entities within the application)
    ├── repositorios                            (Package for performing operations on the database)
    ├── respostas                               (Responsible for handling API responses)
    ├── router                                  (Organizes and manages API routes)
    └── seguranca                               (Responsible for encrypting and decrypting credentials)


```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

The API `.env` file must have the format:
```
DB_USUARIO=user_database
DB_SENHA=password_database
DB_NOME=name_database
API_PORT=listen_port
SECRET_KEY=secret_key
```

The Web Application directories and packages are organized as follows (starting from the root `webapp/`):

```
├── assets
│   ├── css
│   └── js
├── go.mod
├── go.sum
├── main.go                              (Main file)
├── .env 
├── src
│   ├── config                           (Loads environment variables)
│   ├── controllers                      (Handles frontend requests and makes requests to the API or returns templates)
│   ├── cookies                          (Package that contains cookie operations)
│   ├── middlewares                      (Middleware for logging of request and authentication)
│   ├── modelos                          (Defines the data structures of entities within the web application)
│   ├── requisicoes                      (Package responsible for handling authenticated requests to the API directly)
│   ├── respostas                        (Responsible for handling responses to frontend)
│   ├── router                           (Organizes and manages web application routes)
│   └── utils                            (functionalities related to templates)
└── views                                (html files and their templates)
```

The Web Application `.env` file must have the format:

```
API_URL=listen_api_url
PORTA=listen_api_port
HASHKEY=for_cookie
BLOCKKEY=for_cookie

```

### Built With

<p>
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go,javascript,css,html,mysql,jquery,bootstrap" />
  </a>
</p>
I also used https://sweetalert2.github.io/. I didn't find the icon for here hehe)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contact

Wenderson Melo - [Linkedin](https://www.linkedin.com/in/wendersonomelo/) - ds.wendersonmelo@gmail.com

Project Link: [https://github.com/wendersoon/social_networking_app_golang](https://github.com/wendersoon/social_networking_app_golang)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/wendersoon/social_networking_app_golang/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/wendersonomelo/
