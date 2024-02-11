# Auction Website - Teapot 418
This test task was done by team `TEAPOT-418` for `INT20H` hackathon, which involves developing an auction website. The backend stack for this project is Golang, and the frontend stack is Vue.js, Nuxt.js. Below, you will find all the necessary information and descriptions related to the test task.

## Description
The goal of this test task is to create ato develop a web application for an online charity auction that allows users to create auctions, place bids and view active lots. 

> The website is hosted on Digital Ocean and can be reached using next address: http://142.93.172.206/
> 
> The website REST API is public and can be reached using next address: http://142.93.172.206:6969/

**Team**:
- [Roman Skok](https://github.com/romesk) - backend
- [Boryslav Ziubrytskyi](https://github.com/BoryslavGlov) - backend
- [Mykola Balii](https://github.com/Kolia913) - frontend

**This project uses**:
- [GitHub Actions](https://docs.github.com/en/actions)
- [Digital Ocean](https://www.digitalocean.com/)
    - [PostgreSQL](https://www.digitalocean.com/products/managed-databases-postgresql)
    - [Droplet](https://www.digitalocean.com/products/droplets)
- [Docker](https://www.docker.com/), [docker-compose](https://docs.docker.com/compose/)
- [Golang](https://go.dev/) (and various packages)
- [Vue.js](https://vuejs.org/), [nuxt](https://nuxt.com/)
- [JWT tokens](https://jwt.io/)

## Functionality

### Sign-in & Sign-up Pages 
Website includes a sign-in and sign-up page to manage user accounts. Users can create a new account by providing their desired user name, full name, and password. Existing users can sign in using their credentials to access full functioanality.

### Auctions 
- All user can view auctions that take place right now. 
- Logged in users can place their own auctions and make bids for existing ones

### Bidding [Comming soon]
Unfortunately, this part is only partly available on the backend and is implemented via websockets.
In the nearest future you will be able to see fully working version of this part ;)