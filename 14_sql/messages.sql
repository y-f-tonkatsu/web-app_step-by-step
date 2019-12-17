-- create table
CREATE TABLE messages (
    msg_id SERIAL PRIMARY KEY,
    auther VARCHAR(20) NOT NULL,
    message VARCHAR(140) NOT NULL,
);

-- data insert
INSERT INTO messages (auther, message) VALUES 
 ('スピノザ(哲学者)', '人が不可能と思うとき、やりたくないと決めているのだ。')
 ,('ラリー・ウォール', 'あなたたちの多くはプログラマの美徳をよく知っている。もちろんこの三つ、怠惰、短気、傲慢。')
 ;