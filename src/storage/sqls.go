package storage

const (
	GetMovies                = "SELECT ф.Код_Фильма, ф.Название, ф.Год_выпуска, ф.Описание, ф.Рейтинг_кинопоиска, DATE_FORMAT(ф.Дата_добавления_фильма, \"%d-%m-%Y\") AS Дата_добавления_фильма, ф.Дата_последнего_редактирования, ф.Доступно_для_просмотра, ф.Ссылка_на_видео, ф.Ссылка_на_постер, ф.Кол_во_просмотров, ф.Электронная_почта, IF(GROUP_CONCAT(ж.Название SEPARATOR ', ') IS NULL, '', GROUP_CONCAT(ж.Название SEPARATOR ', ')) AS Жанры FROM Фильмы ф LEFT JOIN Фильмы_Жанры фхж ON ф.Код_Фильма = фхж.Код_Фильма LEFT JOIN Жанры ж ON ж.Код_Жанра = фхж.Код_Жанра GROUP BY ф.Код_Фильма ORDER BY ф.Код_Фильма"
	GetMoviesSelect          = "SELECT Код_Фильма AS Код, Название FROM Фильмы WHERE Название LIKE ? LIMIT ? OFFSET ?"
	GetMoviesByFilter        = "CALL get_MoviesByFilterByFilmOrGenre(?, ?)"
	GetMovieWithoutLinkGenre = "SELECT Фильмы.*  FROM Фильмы LEFT JOIN Фильмы_Жанры ON Фильмы.Код_Фильма = Фильмы_Жанры.Код_Фильма WHERE Фильмы_Жанры.Код_Фильма is NULL;"
	GetMovieByID             = "SELECT * FROM Фильмы WHERE Код_Фильма = ?"
	GetLastMovieId           = "SELECT Код_Фильма FROM Фильмы ORDER BY Код_Фильма DESC LIMIT 1"
	GetMoviesCount           = "SELECT COUNT(*) FROM Фильмы WHERE Название LIKE ?"
	UpdateMovie              = "UPDATE Фильмы SET Название = ?, Год_выпуска = ?, Описание = ?, Рейтинг_кинопоиска = ?, Доступно_для_просмотра, Ссылка_на_видео = ?, Ссылка_на_постер = ?, Кол_во_просмотров = ?, Дата_последнего_редактирования = CURRENT_DATE() WHERE Код_Фильма = ?"
	CreateMovie              = "INSERT INTO Фильмы (`Название`, `Год_выпуска`, `Описание`, `Рейтинг_кинопоиска`, `Доступно_для_просмотра`, `Ссылка_на_видео`, `Ссылка_на_постер`, `Кол_во_просмотров`, `Электронная_почта`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	DeleteMovieByMovieID     = "DELETE FROM Фильмы WHERE Код_Фильма = ?"

	GetGenres                 = "SELECT * FROM Жанры"
	GetGenreByID              = "SELECT * FROM Жанры WHERE Код_Жанра = ?"
	GetGenreIdLast            = "SELECT Код_Жанра FROM Жанры ORDER BY Код_Жанра DESC LIMIT 1"
	CreateGenre               = "INSERT INTO Жанры (`Название`, `Описание`) VALUES (?, ?)"
	UpdateGenre               = "UPDATE Жанры SET Название = ?, Описание = ? WHERE Код_Жанра = ?"
	DeleteGenreByGenreID      = "DELETE FROM Жанры WHERE Код_Жанра = ?"
	DeleteMovieGenreByMovieId = "DELETE FROM Фильмы_Жанры WHERE Код_Фильма = ?"

	GetMovieHasGenre               = "SELECT Фильмы.Код_Фильма, Фильмы.Название AS Название_фильма, Жанры.Код_Жанра, Жанры.Название AS Название_жанра FROM Фильмы_Жанры JOIN Фильмы ON Фильмы_Жанры.Код_Фильма = Фильмы.Код_Фильма JOIN Жанры ON Фильмы_Жанры.Код_Жанра = Жанры.Код_Жанра;"
	GetMovieGenreByMovieID         = "CALL get_FilmHasGenresByFilmId(?)"
	CreateMovieHasGenres           = "INSERT INTO Фильмы_Жанры (`Код_Фильма`, `Код_Жанра`) VALUES (?, ?)"
	DeleteMovieGenreByMovieGenreID = "DELETE FROM Фильмы_Жанры WHERE Код_Фильма = ? AND Код_Жанра = ?"

	GetStatisticMovie    = "SELECT COUNT(*) FROM Фильмы WHERE Дата_добавления_фильма >= ? AND Дата_добавления_фильма < ?"
	GetStatisticUsers    = "SELECT COUNT(*) FROM Пользователи WHERE Дата_добавления >= ? AND Дата_добавления < ?"
	GetStatisticComments = "CALL get_StatisticComments(?, ?)"

	GetReportMovie    = "CALL get_FilmAddDiff(?, ?)"
	GetReportUsers    = "SELECT Электронная_почта, ФИО, Логин, Пол, Дата_добавления FROM Пользователи WHERE Дата_добавления > ? AND Дата_добавления < ?;"
	GetReportComments = "SELECT Эл_почта_пользователя, Фильмы.Название, Комментарии.Текст_комментария, Комментарии.Дата_публикации_комментария FROM Комментарии JOIN Фильмы ON Комментарии.Код_Фильма = Фильмы.Код_Фильма WHERE Фильмы.Код_Фильма = ? AND Эл_почта_пользователя = ?"

	GetEmployees               = "SELECT e.Электронная_почта, ФИО, Логин, Пароль, Пол, Адрес, DATE_FORMAT(День_рождения, \"%d-%m-%Y\") AS День_рождения, k.Название AS Компания, k.Код_Компании FROM Сотрудники e JOIN Компании k ON e.Код_Компании = k.Код_Компании"
	GetEmployeeByEmail         = "CALL get_EmployeeByEmail(?)"
	GetEmployeeByEmailPassword = "CALL get_EmployeeByEmailAndPass(?, ?)"
	CreateEmployee             = "INSERT INTO Сотрудники (Электронная_почта, ФИО, Логин, Пароль, Пол, Адрес, День_рождения, Код_Компании) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	UpdateEmployee             = "UPDATE Сотрудники SET Электронная_почта = ?, ФИО = ?, Логин = ?, Пароль = ?, Пол = ?, Адрес = ?, День_рождения = ?, Код_Компании = ? WHERE Электронная_почта = ?"
	DeleteEmployee             = "DELETE FROM Сотрудники WHERE Электронная_почта = ?"

	GetCompanyCount  = "SELECT COUNT(*) FROM Компании WHERE Название LIKE ?"
	GetCompanySelect = "SELECT Код_Компании AS Код, Название FROM Компании WHERE Название LIKE ? LIMIT ? OFFSET ?"

	GetUsers               = "SELECT `Электронная_почта`, `ФИО`, `Логин`, `Пол`, `Код_Компании`, `Адрес`, `День_рождения` FROM Пользователи u JOIN Компании k ON u.Код_Компании = k.Код_Компании"
	GetUserByEmail         = "SELECT Электронная_почта, ФИО, Логин, Пароль, Пол FROM Пользователи WHERE Электронная_почта = ?"
	UpdateUserByEmail      = "UPDATE Пользователи SET Электронная_почта = ?, ФИО = ?, Логин = ?, Пароль = ?, Пол = ? WHERE Электронная_почта = ?"
	GetUsersSelect         = "SELECT Электронная_почта AS Код, ФИО AS Название FROM Пользователи WHERE ФИО LIKE ? OR Электронная_почта LIKE ? LIMIT ? OFFSET ?"
	GetUsersCount          = "SELECT COUNT(*) FROM Пользователи WHERE ФИО LIKE ? OR Электронная_почта LIKE ?"
	GetUserByEmailPassword = "SELECT Электронная_почта, ФИО, Логин, Пароль, Пол FROM Пользователи WHERE Электронная_почта = ? AND Пароль = ?"

	GetCommentById                = "SELECT Код_Комментария, Эл_почта_пользователя, Фильмы.Название, Комментарии.Текст_комментария, Комментарии.Дата_публикации_комментария, IF(Эл_почта_пользователя = ?, 1, 0) AS Создан_текущим_пользователем FROM Комментарии JOIN Фильмы ON Комментарии.Код_Фильма = Фильмы.Код_Фильма WHERE Комментарии.Код_Комментария = ?"
	GetCommentsByMovie            = "SELECT Код_Комментария, Эл_почта_пользователя, Фильмы.Название, Комментарии.Текст_комментария, Комментарии.Дата_публикации_комментария, IF(Эл_почта_пользователя = ?, 1, 0) AS Создан_текущим_пользователем FROM Комментарии JOIN Фильмы ON Комментарии.Код_Фильма = Фильмы.Код_Фильма WHERE Фильмы.Код_Фильма = ? ORDER BY Комментарии.Код_Комментария DESC"
	GetLastCommentByMovieAndEmail = "SELECT Код_Комментария, Эл_почта_пользователя, Фильмы.Название, Комментарии.Текст_комментария, Комментарии.Дата_публикации_комментария, 1 AS Создан_текущим_пользователем FROM Комментарии JOIN Фильмы ON Комментарии.Код_Фильма = Фильмы.Код_Фильма WHERE Фильмы.Код_Фильма = ? AND Эл_почта_пользователя = ? AND Комментарии.Дата_публикации_комментария = CURRENT_DATE() ORDER BY Код_Комментария DESC LIMIT 1"
	CreateComment                 = "INSERT INTO `Комментарии` (`Эл_почта_пользователя`, `Код_Фильма`, `Текст_комментария`, `Дата_публикации_комментария`) VALUES (?, ?, ?, CURRENT_DATE());"
	UpdateComment                 = "UPDATE `Комментарии` SET `Текст_комментария` = ? WHERE `Код_Комментария` = ?"
	DeleteComment                 = "DELETE FROM Комментарии WHERE Код_Комментария = ?"
)
