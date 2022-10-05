package models

type (
	ResponseSelect struct {
		Data          []DataSelect `json:"data"`
		TotalElements uint         `json:"totalElements"`
		Page          uint         `json:"page"`
		Size          uint         `json:"size"`
	}

	Movie struct {
		ID           uint8    `db:"Код_Фильма"`
		Title        string   `db:"Название"`
		Year         uint32   `db:"Год_выпуска"`
		Description  string   `db:"Описание"`
		KpRating     *float32 `db:"Рейтинг_кинопоиска"`
		DateAdd      string   `db:"Дата_добавления_фильма"`
		DateLastEdit string   `db:"Дата_последнего_редактирования"`
		Available    uint     `db:"Доступно_для_просмотра"`
		VideoLink    string   `db:"Ссылка_на_видео"`
		PictureLink  string   `db:"Ссылка_на_постер"`
		CountView    uint8    `db:"Кол_во_просмотров"`
		AddEmpl      string   `db:"Электронная_почта"`
		Genres       *string  `db:"Жанры"`
	}
	MovieHasGenre struct {
		MovieID    uint8  `db:"Код_Фильма"`
		MovieTitle string `db:"Название_фильма"`
		GenreID    uint8  `db:"Код_Жанра"`
		GenreTitle string `db:"Название_жанра"`
	}
	MovieLight struct {
		MovieID    uint8  `db:"Код_Фильма"`
		MovieTitle string `db:"Название_фильма"`
	}
	MovieReport struct {
		MovieTitle   string `db:"Название"`
		MovieDateAdd string `db:"Дата_добавления_фильма"`
	}
	Genre struct {
		ID          uint8  `db:"Код_Жанра"`
		Title       string `db:"Название"`
		Description string `db:"Описание"`
	}
	GenreLight struct {
		GenreID    uint8  `db:"Код_Жанра"`
		GenreTitle string `db:"Название_жанра"`
	}
	CommentStatistic struct {
		MovieTitle string `db:"Название"`
		MovieCount uint8  `db:"Количество_комментариев"`
	}
	UsersReport struct {
		UserEmail     string `db:"Электронная_почта"`
		UserName      string `db:"ФИО"`
		UserLogin     string `db:"Логин"`
		UserFloor     string `db:"Пол"`
		UserDateAdded string `db:"Дата_добавления"`
	}
	Comment struct {
		ID                   string `db:"Код_Комментария"`
		UserEmail            string `db:"Эл_почта_пользователя"`
		MovieTitle           string `db:"Название"`
		CommentText          string `db:"Текст_комментария"`
		CommentDate          string `db:"Дата_публикации_комментария"`
		CreatedByCurrentUser string `db:"Создан_текущим_пользователем"`
	}
	DataSelect struct {
		ID    string `db:"Код"`
		Title string `db:"Название"`
	}
	Employee struct {
		Email     string `db:"Электронная_почта"`
		FIO       string `db:"ФИО"`
		Login     string `db:"Логин"`
		Password  string `db:"Пароль"`
		Floor     string `db:"Пол"`
		Address   string `db:"Адрес"`
		Birthday  string `db:"День_рождения"`
		Company   string `db:"Компания"`
		CompanyID uint8  `db:"Код_Компании"`
	}
	User struct {
		Email    string `db:"Электронная_почта"`
		FIO      string `db:"ФИО"`
		Login    string `db:"Логин"`
		Password string `db:"Пароль"`
		Floor    string `db:"Пол"`
	}
)

const DEFAULT_SELECT_LIMIT = uint(10)
