package config

import (
	"net"
	"time"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/hasher"
	"github.com/milnner/b_modules/models"
)

var (
	Area                           []string
	RootConnString                 string
	AreasObjs                      []models.Area
	Users                          []string
	UsersObjs                      []models.User
	Classes                        []string
	ClassesObjs                    []models.Class
	Contents                       []string
	ContentObjs                    []models.Content
	ClassHasUser                   []string
	ClassHasContent                []string
	UserCanEditeClass              []string
	Activity                       []string
	ImageActivity                  []string
	ImageActivityObjs              []models.ImageActivity
	TextActivity                   []string
	TextActivityObjs               []models.TextActivity
	OneQuestionNAnswerActivity     []string
	OneQuestionNAnswerActivityObjs []models.OneQuestionNAnswerActivity
	AnswerNToOne                   []string
	AnswerNToOneObjs               []models.AnswerNToOne
	DatabaseConn                   *database.DatabaseConn
)
var UserMockPassword = "testetesteteste"

func SetDBData() {

	timeStr := "2020-10-01 15:50:10"
	timeObj, _ := time.Parse(time.DateTime, timeStr)
	Area = []string{
		"INSERT INTO `area`(`id`, `title`, `description`, `owner_id`, `creation_datetime`) VALUES (1,'area1','area1',1, '" + timeStr + "')",
		"INSERT INTO `area`(`id`, `title`, `description`, `owner_id`, `creation_datetime`) VALUES (2,'area2','area2',1, '" + timeStr + "')",
		"INSERT INTO `area`(`id`, `title`, `description`, `owner_id`, `creation_datetime`) VALUES (3,'area3','area3',1, '" + timeStr + "')",
		"INSERT INTO `area`(`id`, `title`, `description`, `owner_id`, `creation_datetime`) VALUES (4,'area4','area4',1, '" + timeStr + "')",
	}

	AreasObjs = []models.Area{
		*models.NewArea(1, "area1", "area1", 1, timeObj, 1),
		*models.NewArea(2, "area2", "area2", 1, timeObj, 1),
		*models.NewArea(3, "area3", "area3", 1, timeObj, 1),
		*models.NewArea(4, "area4", "area4", 1, timeObj, 1),
	}

	hash, _ := hasher.NewBcryptHasher().Hash([]byte(UserMockPassword))

	Users = []string{
		"INSERT INTO `users`(`id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES (1,'name1','surname1','user1@1.com', 0, '" + timeStr + "', '" + timeStr + "', 'male','" + string(hash) + "');",
		"INSERT INTO `users`(`id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES (2,'name2','surname2','user2@2.com', 0, '" + timeStr + "', '" + timeStr + "', 'female','" + string(hash) + "');",
		"INSERT INTO `users`(`id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES (3,'name3','surname3','user3@3.com', 0, '" + timeStr + "', '" + timeStr + "', 'other','" + string(hash) + "');",
		"INSERT INTO `users`(`id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES (4,'name4','surname4','user4@4.com', 0, '" + timeStr + "', '" + timeStr + "', 'male','" + string(hash) + "');",
		"INSERT INTO `users`(`id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES (5,'name5','surname5','user5@5.com', 0, '" + timeStr + "', '" + timeStr + "', 'female','" + string(hash) + "');",
		"INSERT INTO `users`(`id`, `name`, `surname`, `email`, `professor`, `entry_date`, `bourn_date`, `sex`, `hash`) VALUES (6,'name6','surname6','user6@6.com', 0, '" + timeStr + "', '" + timeStr + "', 'other','" + string(hash) + "');",
	}

	UsersObjs = []models.User{
		*models.NewUser(1, "name1", "surname1", "user1@1.com", 0, timeObj, timeObj, "write", "male", string(hash), 1),
		*models.NewUser(2, "name2", "surname2", "user2@2.com", 0, timeObj, timeObj, "write", "female", string(hash), 1),
		*models.NewUser(3, "name3", "surname3", "user3@3.com", 0, timeObj, timeObj, "write", "other", string(hash), 1),
		*models.NewUser(4, "name4", "surname4", "user4@4.com", 0, timeObj, timeObj, "write", "male", string(hash), 1),
		*models.NewUser(5, "name5", "surname5", "user5@5.com", 0, timeObj, timeObj, "write", "female", string(hash), 1),
		*models.NewUser(6, "name6", "surname6", "user6@6.com", 0, timeObj, timeObj, "write", "other", string(hash), 1),
	}

	Classes = []string{
		"INSERT INTO `classes`(`id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`) VALUES (1,'class1','class1desc','" + timeStr + "',1, '" + timeStr + "', 1)",
		"INSERT INTO `classes`(`id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`) VALUES (2,'class2','class2desc','" + timeStr + "',2, '" + timeStr + "', 1)",
		"INSERT INTO `classes`(`id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`) VALUES (3,'class3','class3desc','" + timeStr + "',3, '" + timeStr + "', 1)",
		"INSERT INTO `classes`(`id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`) VALUES (4,'class4','class4desc','" + timeStr + "',1, '" + timeStr + "', 1)",
		"INSERT INTO `classes`(`id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`) VALUES (5,'class5','class5desc','" + timeStr + "',2, '" + timeStr + "', 1)",
		"INSERT INTO `classes`(`id`, `title`, `description`, `creation_datetime`, `creator_user_id`, `last_update`, `area_id`) VALUES (6,'class6','class6desc','" + timeStr + "',3, '" + timeStr + "', 1)",
	}
	ClassesObjs = []models.Class{
		*models.NewClass(1, "class1", "class1desc", timeObj, 1, 1, timeObj, 1),
		*models.NewClass(2, "class2", "class2desc", timeObj, 2, 1, timeObj, 1),
		*models.NewClass(3, "class3", "class3desc", timeObj, 3, 1, timeObj, 1),
		*models.NewClass(4, "class4", "class4desc", timeObj, 1, 1, timeObj, 1),
		*models.NewClass(5, "class5", "class5desc", timeObj, 2, 1, timeObj, 1),
		*models.NewClass(6, "class6", "class6desc", timeObj, 3, 1, timeObj, 1),
	}
	Contents = []string{
		"INSERT INTO `contents`(`id`,  `title`, `description`, `creation_datetime`, `last_update`, `area_id`) VALUES (1, 'title1', 'description1', '" + timeStr + "','" + timeStr + "',1)",
		"INSERT INTO `contents`(`id`, `title`, `description`, `creation_datetime`, `last_update`, `area_id`) VALUES (2, 'title2', 'description2', '" + timeStr + "','" + timeStr + "',1)",
		"INSERT INTO `contents`(`id`, `title`, `description`, `creation_datetime`, `last_update`, `area_id`) VALUES (3, 'title3', 'description3', '" + timeStr + "','" + timeStr + "',1)",
	}

	ContentObjs = []models.Content{
		*models.NewContent(1, timeObj, "title1", "description1", timeObj, 1, 1),
		*models.NewContent(2, timeObj, "title2", "description2", timeObj, 1, 1),
		*models.NewContent(3, timeObj, "title3", "description3", timeObj, 1, 1),
	}
	ClassHasUser = []string{
		"INSERT INTO `class_has_user`(`id`, `entry_date`, `user_id`, `class_id`) VALUES (1,'" + timeStr + "',1,1)",
		"INSERT INTO `class_has_user`(`id`, `entry_date`, `user_id`, `class_id`) VALUES (2,'" + timeStr + "',2,1)",
		"INSERT INTO `class_has_user`(`id`, `entry_date`, `user_id`, `class_id`) VALUES (3,'" + timeStr + "',3,1)",
		"INSERT INTO `class_has_user`(`id`, `entry_date`, `user_id`, `class_id`) VALUES (4,'" + timeStr + "',2,3)",
	}
	ClassHasContent = []string{
		"INSERT INTO `class_has_content`(`id`, `class_id`, `content_id`, `position`) VALUES (1,1,1,0)",
		"INSERT INTO `class_has_content`(`id`, `class_id`, `content_id`, `position`) VALUES (2,1,2,1)",
		"INSERT INTO `class_has_content`(`id`, `class_id`, `content_id`, `position`) VALUES (3,1,3,2)",
		"INSERT INTO `class_has_content`(`id`, `class_id`, `content_id`, `position`) VALUES (4,2,3,0)",
	}

	UserCanEditeClass = []string{
		"INSERT INTO `user_can_edit_class`(`id`, `editor_user_id`, `class_id`, `entry_date`) VALUES (1,1,1, '" + timeStr + "')",
		"INSERT INTO `user_can_edit_class`(`id`, `editor_user_id`, `class_id`, `entry_date`) VALUES (2,2,1, '" + timeStr + "')",
		"INSERT INTO `user_can_edit_class`(`id`, `editor_user_id`, `class_id`, `entry_date`) VALUES (3,3,1, '" + timeStr + "')",
		"INSERT INTO `user_can_edit_class`(`id`, `editor_user_id`, `class_id`, `entry_date`) VALUES (4,2,3, '" + timeStr + "')",
	}

	imgStr := "0x10101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010"
	img := []byte(imgStr)

	ImageActivity = []string{
		"INSERT INTO `image_activities`(`id`, `area_id`, `title`, `_blob`, `last_update`) VALUES (1, 1, 'image1', '" + imgStr + "' ,'" + timeStr + "')",
		"INSERT INTO `image_activities`(`id`, `area_id`, `title`, `_blob`, `last_update`) VALUES (2, 1,'image2', '" + imgStr + "' ,'" + timeStr + "')",
	}
	ImageActivityObjs = []models.ImageActivity{
		*models.NewImageActivity(1, 1, "image1", img, timeObj, 1),
		*models.NewImageActivity(2, 1, "image2", img, timeObj, 1),
	}

	textStr := "Opa olha eu aqui.,Opa olha eu aquiOpa olha eu aquiOpa olha eu aquiOpa olha eu aquiOpa olha eu aquiOpa olha eu aquiOpa olha eu aquiOpa olha eu aqui"
	textBlob := []byte(textStr)

	TextActivity = []string{
		"INSERT INTO `text_activities`(`id`, `area_id`, `title`, `_blob`, `last_update`) VALUES (1, 1,'text1','" + textStr + "','" + timeStr + "')",
		"INSERT INTO `text_activities`(`id`, `area_id`, `title`, `_blob`, `last_update`) VALUES (2, 1,'text2','" + textStr + "','" + timeStr + "')",
	}
	TextActivityObjs = []models.TextActivity{
		*models.NewTextActivity(1, 1, "text1", textBlob, timeObj, 1),
		*models.NewTextActivity(2, 1, "text2", textBlob, timeObj, 1),
	}

	question1Txt := "question 1"
	question1 := []byte(question1Txt)
	question2Txt := "question 2"
	question2 := []byte(question2Txt)

	OneQuestionNAnswerActivity = []string{
		"INSERT INTO `one_question_n_answer_activities`(`id`, `area_id`, `question`, `last_update`, `activated`) VALUES (1,1,'" + question1Txt + "','" + timeStr + "',1)",
		"INSERT INTO `one_question_n_answer_activities`(`id`, `area_id`, `question`, `last_update`, `activated`) VALUES (2,1,'" + question2Txt + "','" + timeStr + "',1)",
	}

	OneQuestionNAnswerActivityObjs = []models.OneQuestionNAnswerActivity{
		*models.NewOneQuestionNAnswerActivity(1, 1, question1, timeObj, 1),
		*models.NewOneQuestionNAnswerActivity(2, 1, question2, timeObj, 1),
	}

	answer1Txt := "Resposta1"
	answer1 := []byte(answer1Txt)

	AnswerNToOne = []string{
		"INSERT INTO `answer_n_to_one`(`id`, `area_id`, `one_question_n_answer_activity_id`, `correctness`, `answer`, `activated`) VALUES (1,1,1,100,'" + answer1Txt + "',1)",
		"INSERT INTO `answer_n_to_one`(`id`, `area_id`, `one_question_n_answer_activity_id`, `correctness`, `answer`, `activated`) VALUES (2,1,1,0,'" + answer1Txt + "',1)",
	}

	AnswerNToOneObjs = []models.AnswerNToOne{
		*models.NewAnswerNToOne(1, 1, 1, 100, answer1, 1),
		*models.NewAnswerNToOne(2, 1, 1, 0, answer1, 1),
	}

}

func SetRootDatabaseConn() {
	rootConnString := "root:root@tcp(127.0.0.1:3306)/ardeo"

	port := "3306"
	DatabaseConn = database.NewDatabaseConn()

	database.SetRoot(DatabaseConn, rootConnString)

	target := "127.0.0.1:" + port
	conn, err := net.DialTimeout("tcp", target, 10*time.Second)
	if err != nil {
		panic(err)
	}
	conn.Close()
}
