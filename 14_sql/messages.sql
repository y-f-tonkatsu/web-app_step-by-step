-- create table
CREATE TABLE messages (
    id INT(11) AUTO_INCREMENT NOT NULL, 
    name VARCHAR(30) ,
    message VARCHAR(30) ,
    PRIMARY KEY (id)
);

-- data insert
INSERT INTO messages (name, message) VALUES 
 ('スピノザ(哲学者)', '人が不可能と思うとき、やりたくないと決めているのだ。')
 ,('ラリー・ウォール', 'あなたたちの多くはプログラマの美徳をよく知っている。もちろんこの三つ、怠惰、短気、傲慢。')
 ;