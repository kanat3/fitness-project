/*=============== SHOW MENU ===============*/
const navMenu = document.getElementById('nav-menu'),
      navToggle = document.getElementById('nav-toggle'),
      navClose = document.getElementById('nav-close')

/* Menu show */
navToggle.addEventListener('click', () =>{
   navMenu.classList.add('show-menu')
})

/* Menu hidden */
navClose.addEventListener('click', () =>{
   navMenu.classList.remove('show-menu')
})

/*=============== SEARCH ===============*/
const search = document.getElementById('search'),
      searchBtn = document.getElementById('search-btn'),
      searchClose = document.getElementById('search-close')

/* Search show */
searchBtn.addEventListener('click', () =>{
   search.classList.add('show-search')
})

/* Search hidden */
searchClose.addEventListener('click', () =>{
   search.classList.remove('show-search')
})

/*=============== LOGIN ===============*/
const login = document.getElementById('login'),
      loginBtn = document.getElementById('login-btn'),
      loginClose = document.getElementById('login-close')

/* Login show */
loginBtn.addEventListener('click', () =>{
   login.classList.add('show-login')
})

/* Login hidden */
loginClose.addEventListener('click', () =>{
   login.classList.remove('show-login')
})

const button = document.getElementById('button');
const menu = document.getElementById('menu');

button.addEventListener('click', () =>{
    if(button.dataset.trigger == 'false'){
        button.innerText = 'danya_sidorov@mail.ru'
        menu.style.display = 'block';
        button.dataset.trigger = true;
    }else{
        button.innerText = button.dataset.text;
        menu.style.display = 'none';
        button.dataset.trigger = false;
    }
});

const button2 = document.getElementById('button2');
const menu2 = document.getElementById('menu2');

button2.addEventListener('click', () =>{
    if(button2.dataset.trigger == 'false'){
        button2.innerText = 'ПОНЕДЕЛЬНИК. Завтрак: Сырники, смузи из свежих ягод. Первый перекус: 30 г грецких орехов. Обед: Куриный бульон, овощной салат. Второй перекус: Зеленое яблоко. Ужин: Отварная говядина со свежим салатом. ВТОРНИК.Завтрак: Овсяная каша с грушами в сиропе без сахара. Первый перекус: Натуральный йогурт. Обед: Рагу из овощей, борщ без картофеля. Второй перекус: Протеиновый коктейль. Ужин: Тушеная капуста, отварная куриная грудка. СРЕДА. Завтрак: Гречневая каша с молоком. Первый перекус: Смузи из яблок. Обед: Грибной крем-суп, мясной салат. Второй перекус: Груша. Ужин: Рагу из говядины с овощами. ЧЕТВЕРГ. Завтрак: Творожная запеканка, яблоко. Первый перекус: Любой фрукт кроме запрещенных. Обед: Суп с фрикадельками без картофеля, овощной салат. Второй перекус: Протеиновый коктейль. Ужин: Минтай на пару со свежими овощами. ПЯТНИЦА. Завтрак: Геркулес, творожный сыр. Первый перекус: Яблоко или груша. Обед: Говяжий бульон, салат из овощей и брынзы. Второй перекус: Натуральный йогурт. Ужин: Мясная запеканка, овощной салат. ВЫХОДНЫЕ ДНИ. Завтрак: Рисовая каша или лапша на молоке. Первый перекус: Любой фрукт на выбор. Обед: Отварная говядина, овощной бульон. Второй перекус: 30 г фундука. Ужин: Творожная запеканка с отварной курицей.';
        menu2.style.display = 'block';
        button2.dataset.trigger = true;
    }else{
        button2.innerText = button2.dataset.text;
        menu2.style.display = 'none';
        button2.dataset.trigger = false;
    }
});


