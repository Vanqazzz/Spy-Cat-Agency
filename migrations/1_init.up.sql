CREATE TABLE cats (
    id SERIAL PRIMARY KEY,           
    name TEXT NOT NULL UNIQUE,      
    years_of_exp INTEGER NOT NULL,   
    breed TEXT NOT NULL,           
    salary INTEGER NOT NULL          
);