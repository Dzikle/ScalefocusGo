CREATE TABLE stories(
storyid INT PRIMARY KEY,
title TEXT NOT NULL,
score INT NOT NULL,
timestamp DATETIME NOT NULL
);

-- name: GetStories :many
SELECT storyid, title, score 
FROM stories s;

-- name: SaveStories :execresult
INSERT INTO stories (storyid,title,score,timestamp) 
values(?,?,?,?);

-- name: GetLastStoryTimeStamp :one
SELECT s.timestamp FROM stories s 
ORDER BY s.timestamp DESC LIMIT 1;