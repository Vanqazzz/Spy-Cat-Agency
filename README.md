# Spy-Cat-Agency

Spy-Cat-Agency crud operations app with possibility to create,update,remove,update:

Task requirements:
- Breed validator done
- Gin default logger done

Not done: 
- Ability to update mission targets
  - Ability to mark it as completed
  
- Ability to update Notes
  - Notes cannot be updated if either the target or the mission is completed

- Ability to delete targets from an existing mission
  - A target cannot be deleted if it is already completed

- Ability to add targets to an existing mission


Endpoints (Tested by Postman)

Create Cat:
POST /createcat
Body: { "name": "Whiskers", "age": 5, "breed": "Siamese", "salary": 500 }

Create Mission:
POST /createmission
Body:  "Cat_id": 1, "Target_id": 1, "complete_state": false

Create Target:
POST"/createtarget" Body: { "Target_name": "Name", "country": "Country", "notes": "Notes",  "complete_state_target": false
}

Get List Cats:
GET /listallcats (without body)

Get SingleCat: 
GET/getasinglecat Body: {"Name": "Amon"}

Get SingleMission:
GET/singlemission"
Body: {"Id": "1"}

Get Allmissions:
GET/allmissions", (without body)

Delete Cat:
DELETE"/delete", Body: {"Name": "Name"}

Delete Missions:
DELETE"/deletemission",  Body: {"Id": 1}

Assign cat:
PUT"/assigncat",  Body: {"Cat_id": 1, "Id": 2}

Update Cat salary:
PUT/update", Body: {"Name": "cat", "salary" 1000}

Update Mission:
PUT(/missionupdate") , {"Id": 1, "complemete_state" true}


## Stack
Language: Go

Framework: Gin

Database: PostgreSQL

### Instalattion

Clone repository in your folder
```cmd
git clone https://github.com/Vanqazzz/Spy-Cat-Agency
cd Spy-Cat-Agency
```

## Build container
```
docker-compose up --build
```

