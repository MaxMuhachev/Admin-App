
//Удаление фильма
$('.deleteMovie').click(function () {
  let movieId = this.value;
  let rowToRemove = $(this).parent().parent();
  $.ajax({
    method: "POST",
    url: "/movies/delete",
    dataType: 'json',
    async: true,
    data: {id: movieId},
    success: function (response) {
      rowToRemove.remove();
      toast({heading: "", text: "Фильм удален", type: "success"});
    },
    error: function (response) {
      toast({heading: "", text: "Ошибка удаленмя фильма" + response.responseText, type: "error"});
    }
  })
});

//Удаление пользователя
$('.deleteEmployee').click(function () {
  let movieId = this.value;
  let rowToRemove = $(this).parent().parent();
  $.ajax({
    method: "POST",
    url: "/employees/delete",
    dataType: 'json',
    async: true,
    data: {id: movieId},
    success: function (response) {
      rowToRemove.remove();
      toast({heading: "", text: "Сотрудник удален", type: "success"});
    },
    error: function (response) {
      toast({heading: "", text: "Ошибка удаленмя сотрудника" + response.responseText, type: "error"});
    }
  })
});

//Удаление жанра
$('.deleteGenre').click(function () {
  let movieId = this.value;
  let rowToRemove = $(this).parent().parent();
  $.ajax({
    method: "POST",
    url: "/genres/delete",
    dataType: 'json',
    async: true,
    data: {id: movieId},
    success: function (response) {
      rowToRemove.remove();
      toast({heading: "", text: "Жанр удален", type: "success"});
    },
    error: function (response) {
      rowToRemove.remove();
      toast({heading: "", text: "Ошибка удаленмя жанра" + response.responseText, type: "error"});
    }
  })
});

function toast(toastData) {
  var config = {
    loader: false,
    showHideTransition: 'fade',
    position:  'bottom-right',
    hideAfter: true,
    heading: toastData.heading,
    text: toastData.text,
    icon: toastData.type
  };
  $.toast(config)
}