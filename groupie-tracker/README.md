# Groupie Tracker
Project Overview
---
Groupie Tracker is a web application that utilizes data from a provided API to showcase information about various bands and artists. The project focuses on creating a user-friendly website that displays detailed information about the artists, including their members, concert locations, and dates. The application also implements an event-driven feature that triggers a server call, enabling dynamic interaction between the client and server.

API Structure
---
The API provided consists of four key components:

Artists: Contains details about bands and artists, including their name(s), image, the year they began their activity, the release date of their first album, and the band members.
Locations: Provides information on the upcoming concert locations for each artist.
Dates: Contains the dates of the upcoming concerts.
Relation: Links the artists with their respective concert dates and locations.

Features
---
`Data Visualization:` The website displays the data from the API using cards, lists. The design and structure are created with user experience in mind, making the information accessible and engaging.

`Client-Server Interaction:` A key feature of the project is the implementation of a client-server interaction, where an event/action triggered by the client initiates a server request. The server then responds with the required data, which is dynamically displayed on the website.

# Getting Started
Prerequisites
---
Go 1.22.3 or higher

Basic knowledge of HTML, CSS, and Go

Access to the provided API

Installation
---

Clone the repository:

```
git clone https://learn.zone01oujda.ma/git/mfir/groupie-tracker.git
```

```
cd groupie-tracker/server
```

Install dependencies (if any):

```
go mod web
```

Run the server:

```
go run main.go
```

Open your browser and navigate to `http://localhost:8081` to view the website.

Directory Structure
---
main.go: The main entry point of the application.

templates/: Contains HTML templates for the website.

static/: Contains static files such as CSS and images.

handlers/ and methods/: Contains the server-side code for handling client requests.

Usage
---
Once the server is running, you can navigate through the website to explore various artists, concert locations, and dates. The client-server interaction allows you to trigger specific actions, such as fetching updated concert dates or locations, which are then displayed on the site.

Contributors
---
<a href="https://learn.zone01oujda.ma/git/amazighi">
  <img src= "https://avatars.githubusercontent.com/u/125706060?size=40" title="Abdessamad Mazighi"/>
</a>
<a href="https://learn.zone01oujda.ma/git/rachnit">
  <img src="https://learn.zone01oujda.ma/git/avatars/36d91fe5b20bba8044b652425b93d4d9?size=870" title="Reda Achnit" width="40"height="40"/>
</a><a href="https://learn.zone01oujda.ma/git/mfir">
  <img src="https://learn.zone01oujda.ma/git/avatars/df5adee604dd5c0dca31fd1aa128041d?size=870" title="Mohamed Fri" width="40"height="40"/>
</a>
