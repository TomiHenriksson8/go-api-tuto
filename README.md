# Go API -projekti

## Projektin yleiskuvaus
Go-ohjelmointikielellä rakennettu REST API, joka käyttää MongoDB:tä tietojen tallennukseen. Projekti sisältää perus CRUD-operaatioita, autentikointiin ja valtuutukseen liittyvää middlewarea sekä dokumentaatiota Swaggerin avulla.

## Suoritetut vaiheet

### Vaihe 1. Tutustu tutoriaaleihin
Aloitin perehtymällä Go:n REST API -tutoriaaleihin, joiden avulla opin rakentamaan RESTful-palveluja Go:lla.

### Vaihe 2. API:n toteutus
Toteutin API:n Go-ohjelmointikielellä käyttäen MongoDB:tä tietojen tallennukseen. API:ssa on toteutettu perus CRUD-toiminnot:

- **GET**: Hakee tietoa tietokannasta
- **POST**: Lisää tietoa tietokantaan
- **PUT**: Päivittää olemassa olevaa tietoa
- **DELETE**: Poistaa tietoa tietokannasta

### Vaihe 3. Edistyneet ominaisuudet
Toteutin API:hin seuraavia edistyneitä ominaisuuksia:

- **Autentikointi**: Käyttäjän sisäänkirjautuminen ja JWT-tunnusten luominen.
- **Valtuutus**: Middleware, joka tarkistaa käyttäjän oikeudet reittikohtaisesti.
- **Validointi**: Pyyntöjen validointi API:ssa ennen tietokantaan tallentamista.

### Vaihe 4. Dokumentaatio
API-dokumentaatio on luotu Swaggerin avulla. Dokumentaatio sisältää kaikki API-reitit ja niiden käyttötavat.

- [Swagger-dokumentaatio](https://mygoapiwebapp.azurewebsites.net/api/swagger/index.html)

### Vaihe 5. Julkaisu
API on julkaistu Microsoft Azuren pilvipalvelussa Dockerin avulla.

## Käytetyt teknologiat
- **Kieli**: Go
- **Tietokanta**: MongoDB
- **Pilvi**: Microsoft Azure
- **Konttiteknologia**: Docker
- **Dokumentaatio**: Swagger
