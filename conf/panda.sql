CREATE TABLE `Comment` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Content` text NOT NULL,
  `CommentTime` datetime NOT NULL,
  `CommentType` int(11) NOT NULL,
  `User` varchar(128) NOT NULL DEFAULT 'admin',
  `IP` varchar(64) DEFAULT NULL,
  `UA` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `Portrait` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(128) NOT NULL,
  `Path` varchar(256) NOT NULL,
  `Time` datetime NOT NULL,
  `User` varchar(128) DEFAULT 'NULL',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `User` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(128) NOT NULL,
  `Password` varchar(128) DEFAULT NULL,
  `Permission` int(11),
  `RegisterTime` datetime NOT NULL,
  `LoginTime` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE Move(
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    Content text NOT NULL,
    ImgPath varchar(256) NOT NULL,
    User varchar(128) NOT NULL,
    Time datetime NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `MCR` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Mid` int(11) NOT NULL,
  `Cid` int(11) NOT NULL,
  `Time` datetime NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `Announce` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Content` text NOT NULL,
  `Display` tinyint(1) NOT NULL DEFAULT 0,
  `Prior` int(11) DEFAULT 0,
  `User` varchar(128) NOT NULL,
  `Time` datetime NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


