CREATE TABLE
     IF NOT EXISTS users (
          id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
          name TEXT NOT NULL,
          email TEXT UNIQUE,
          password TEXT,
          created_at TIMESTAMP
          WITH
               TIME ZONE DEFAULT NOW (),
               archived_at TIMESTAMP
          WITH
               TIME ZONE
     );

CREATE TABLE
     IF NOT EXISTS sessions (
          id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
          user_id UUID REFERENCES users (id) NOT NULL,
          session_token TEXT NOT NULL,
          expiry_time TIMESTAMP
          WITH
               TIME ZONE,
               created_at TIMESTAMP
          WITH
               TIME ZONE DEFAULT NOW ()
     );

CREATE TABLE
     IF NOT EXISTS notes (
          id UUID KEY DEFAULT get_random_uuid (),
          FOREIGN KEY (user_id) REFERENCES users (id),
          title TEXT NOT NULL,
          text TEXT,
          created_at TIMESTAMP
          WITH
               TIME ZONE DEFAULT NOW (),
               archived_at TIMESTAMP
          WITH
               TIME ZONE
     )