CREATE TABLE User (
    UserID              INTEGER PRIMARY KEY AUTOINCREMENT,
    Username            TEXT NOT NULL UNIQUE,
    FirstName           TEXT,
    LastName            TEXT,
    Email               TEXT NOT NULL UNIQUE,
    Password            TEXT NOT NULL,
    SessionToken        TEXT,
    SessionExpiration   DATETIME
);

CREATE TABLE Category (
    CategoryID      INTEGER PRIMARY KEY AUTOINCREMENT,
    CategoryName    TEXT NOT NULL UNIQUE
);

CREATE TABLE Post (
    PostID          INTEGER PRIMARY KEY AUTOINCREMENT,
    Title           TEXT NOT NULL,
    Content         TEXT NOT NULL,
    CreatedDate     DATETIME DEFAULT CURRENT_TIMESTAMP,
    UserID          INTEGER,
    -- CategoryID      INTEGER,
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID)
);

CREATE TABLE Comment (
    CommentID       INTEGER PRIMARY KEY AUTOINCREMENT,
    Content         TEXT NOT NULL,
    CreatedDate     DATETIME DEFAULT CURRENT_TIMESTAMP,
    PostID          INTEGER,
    UserID          INTEGER,
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

CREATE TABLE PostCategory (
    PostCategoryID  INTEGER PRIMARY KEY AUTOINCREMENT,
    PostID          INTEGER,
    CategoryID      INTEGER,
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID)
);

CREATE TABLE Interaction (
    InteractionID   INTEGER PRIMARY KEY AUTOINCREMENT,
    PostID          INTEGER,
    UserID          INTEGER,
    Kind            INTEGER NOT NULL CHECK (Kind IN (0, 1)), -- 1 for like, 0 for dislike
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);
