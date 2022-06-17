DROP TABLE IF EXISTS plans;

CREATE TABLE IF NOT EXISTS plans (
  ID              BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	ShopName        VARCHAR(40) NOT NULL,
	MeetPlace       VARCHAR(40) NOT NULL,
	MaxPeopleNumber INT(11),
	MinPeopleNumber INT(11) DEFAULT 1,
	MeetTime        DATETIME NOT NULL,
	PlanStatus      INT(40) DEFAULT 0,
	OwnerUserId     BIGINT(20) UNSIGNED NOT NULL,
	ParticipantUsersCount INT(20) DEFAULT 1
);
