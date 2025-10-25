CREATE TABLE cats (
    id SERIAL PRIMARY KEY,           
    name TEXT NOT NULL UNIQUE,      
    years_of_exp INTEGER NOT NULL,   
    breed TEXT NOT NULL,           
    salary INTEGER NOT NULL          
);

CREATE TABLE target (
    id SERIAL PRIMARY KEY,           
    name TEXT NOT NULL UNIQUE,      
    country TEXT NOT NULL,   
    Notes TEXT,           
    complete_state bool          
);

CREATE TABLE missions (
    id SERIAL PRIMARY KEY,           
    cat_id INTEGER NOT NULL,      
    target_id INTEGER  NOT NULL,       
    complete_state bool,  
    FOREIGN KEY (cat_id) REFERENCES cats(id),
    FOREIGN KEY (target_id) REFERENCES target(id)                  
);

