package wiki

import (
	"testing"
)

func TestParse_1_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 1, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `1 декабря — 335-й день года (336-й в високосные годы) в григорианском календаре.
До конца года остаётся 30 дней.
В XX и XXI веках соответствует 18 ноября юлианского календаря.


== Праздники и памятные дни ==


=== Международные ===
 ООН — Всемирный день борьбы со СПИДом 


=== Национальные ===
 Казахстан — День первого президента Казахстана.
 Каракалпакстан — День каракалпакского языка.
 Россия — День воинской славы России в честь победы русской эскадры под командованием П. С. Нахимова над турецкой эскадрой у мыса Синоп (1853 год). На самом деле сражение произошло 18 (30) ноября 1853 года.
 Румыния — День объединения Румынии (Национальный день).
 Португалия — День независимости.
 Украина — День работников прокуратуры


=== Религиозные ===
 В православной церквипамять мученика Платона (302 или 306 год);
память мучеников Романа диакона и отрока Варула (303 год);
память мучеников Закхея, диакона Гадаринского, и Алфея, чтеца Кесарийского (303 год);
память святого Николая Виноградова, исповедника, пресвитера (1948 год);
Собор святых Эстонской земли.


=== Именины ===
Платон, Роман

`
	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Международные_
- ООН — Всемирный день борьбы со СПИДом

_Национальные_
- Казахстан — День первого президента Казахстана
- Каракалпакстан — День каракалпакского языка
- Россия — День воинской славы России в честь победы русской эскадры под командованием П. С. Нахимова над турецкой эскадрой у мыса Синоп (1853 год). На самом деле сражение произошло 18 (30) ноября 1853 года
- Румыния — День объединения Румынии (Национальный день)
- Португалия — День независимости
- Украина — День работников прокуратуры

_Религиозные_
- Собор святых Эстонской земли (правосл.)

_Именины_
- Платон, Роман
`
	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_4_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 4, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `4 декабря — 338-й день года (339-й в високосные годы) в григорианском календаре.
До конца года остаётся 27 дней.
В XX и XXI веках соответствует 21 ноября юлианского календаря.


== Праздники и памятные дни ==


=== Национальные ===
 Мексика — День ремесленников.
 Россия — День российской информатики.
 Тонга — День провозглашения государства.


=== Религиозные ===
 В католической церквиАнно II
 В православной церквиВведение во храм Пресвятой Богородицы


=== Именины ===
Мария.


== Приметы ==
«Введенье, Ворота зимы». Народ приметил, что в это время бывают морозы:

«Введение накладывает на воду ледение»,
«На Введение — толстое леденье»,
«Введение пришло — зиму привело»,
«Введенские морозы зиму на ум наставляют»,
«Введенские морозы рукавицы на мужика надели, стужу установили, зиму на ум наставили».Но когда на Введение оттепель, то говорили:

«Введение ломает леденье»,
«Если со Введения ляжет глубокая зима — готовь глубокие закрома: будет богатый урожай хлебов»,
«Во Введение мороз — все праздники морозны, тепло — все праздники теплы».На Введение делались пробные выезды на санях, право начинать эти гулянья отводилось молодожёнам. Обряд назывался «казать молодую». В этот день открывались Введенские ярмарки, торги.


== См. также ==


== Примечания ==`
		report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Национальные_
- Мексика — День ремесленников
- Россия — День российской информатики
- Тонга — День провозглашения государства

_Религиозные_
- Введение во храм Пресвятой Богородицы (правосл.)
- Анно II

_Именины_
- Мария

*Приметы*

_«Введенье, Ворота зимы»_
Народ приметил, что в это время бывают морозы:
«Введение накладывает на воду ледение»,
«На Введение — толстое леденье»,
«Введение пришло — зиму привело»,
«Введенские морозы зиму на ум наставляют»,
«Введенские морозы рукавицы на мужика надели, стужу установили, зиму на ум наставили»
Но когда на Введение оттепель, то говорили:
«Введение ломает леденье»,
«Если со Введения ляжет глубокая зима — готовь глубокие закрома: будет богатый урожай хлебов»,
«Во Введение мороз — все праздники морозны, тепло — все праздники теплы»
На Введение делались пробные выезды на санях, право начинать эти гулянья отводилось молодожёнам
Обряд назывался «казать молодую»
В этот день открывались Введенские ярмарки, торги
`
	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_7_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 7, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `7 декабря — 341-й день года (342-й в високосные годы) в григорианском календаре.
До конца года остаётся 24 дня.
В XX и XXI веках соответствует 24 ноября юлианского календаря.


== Праздники и памятные дни ==
См. также: Категория:Праздники 7 декабря


=== Международные ===
 Международный день гражданской авиации.


=== Национальные ===
 США — годовщина нападения на Пёрл-Харбор.
 Украина — День местного самоуправления.


=== Религиозные ===
 Католицизм
 — память святого Амвросия Медиоланского;
 — память Аниана, епископа Шартра;
 — память Юмбера (1148);
 — память Марии (в крещении Бенедетты) Росселло (1880);
 — память Виктора, епископа Пьяченца (375);
 — память святого Сервуса;
 — память святых Поликарпа и Теодора (Феодора) Антиохийских.
 Православие
 — память преподобной Мастридии;
 — память великомученика Меркурия (III век);
 — память великомученицы Екатерины (305-313 годы);
 — память мученицы Августы, мучеников Порфирия Стратилата Александрийского и 200 воинов (305-313 годы);
 — память мученика Меркурия Смоленского (1238 год);
 — память преподобного Меркурия, постника Печерского, в Дальних пещерах (XIV век);
 — память преподобного Симона Сойгинского (1562 год);
 — память священномученика Евграфа Еварестова, пресвитера (1919 год);
 — память священномучеников Евгения Яковлева и Михаила Богородицкого, пресвитеров (1937 год);
 — память священномучеников Александра Левицкого, Алексия Тютюнова, Иоанна Никольского, Корнилия Удиловича и Митрофана Корницкого, пресвитеров (1937 год).


=== Именины ===
Католические: Амвросий, Аниан, Виктор, Мария, Поликарп, Серв(ус), Теодор, Юмбер
Православные: Августа, Александр, Алексей, Григорий, Евгений, Евграф, Екатерина, Ермоген, Иван, Корнелий и Корнилий, Марк, Мастридия, Меркурий, Митрофан, Михаил, Порфирий, Прокопий, Симон, Филофея, Филумен, Христофор
`
	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Международные_
- Международный день гражданской авиации

_Национальные_
- США — годовщина нападения на Пёрл-Харбор
- Украина — День местного самоуправления

_Именины_
- Католические: Амвросий, Аниан, Виктор, Мария, Поликарп, Серв(ус), Теодор, Юмбер
- Православные: Августа, Александр, Алексей, Григорий, Евгений, Евграф, Екатерина, Ермоген, Иван, Корнелий и Корнилий, Марк, Мастридия, Меркурий, Митрофан, Михаил, Порфирий, Прокопий, Симон, Филофея, Филумен, Христофор
`

	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_10_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 10, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `10 декабря — 344-й день года (345-й в високосные годы) в григорианском календаре.
До конца года остаётся 21 день.
В XX и XXI веках соответствует 27 ноября юлианского календаря.


== Праздники и памятные дни ==
См. также: Категория:Праздники 10 декабря


=== Международные ===
 День прав человека
 Всемирный день футбола.


=== Национальные ===
 Россия
 Марий Эл — День марийской письменности.
 Таиланд — День Конституции.


=== Религиозные ===

 Православиепразднование иконы Божией Матери, именуемой «Знамение»;
память знамения Пресвятой Богородицы, бывшее в Новгороде Великом в 1170 году;
память великомученика Иакова Персянина (421);
память преподобного Палладия Александрийского (VI—VII в.);
память святителя Иакова Ростовского, епископа (1392);
обретение мощей благоверного князя Всеволода, во Святом Крещении Гавриила, Новгородского, Псковского чудотворца (1192);
память блаженного Андрея Симбирского (Огородникова) (1841);
Собор новомучеников и исповедников Радонежских;
память преподобномучеников монахов 17-ти в Индии (IV в.);
память преподобного Романа Антиохийского (Сирийского) (V в.);
память священномучеников Николая (Добронравова), архиепископа Владимирского, Василия Соколова, Бориса Ивановского, Феодора Дорофеева, Николая Андреева, Алексия Сперанского, Иоанна Глазкова, Сергия Аманова, Иоанна Хрусталева, Сергия Бредникова, Николая Покровского, Димитрия Беляева, Владимира Смирнова, Иоанна Смирнова, пресвитеров, преподобномучеников Иоасафа (Боева), Кронида (Любимова), архимандритов, Николая (Салтыкова), игумена, Ксенофонта (Бондаренко), иеромонаха, Алексия (Гаврина), монаха, Аполлоса (Федосеева), иеромонаха, Серафима (Крестьянинова), игумена, Никона (Беляева), архимандрита и мученика Иоанна Емельянова (1937);
празднование Курской-Коренной иконы Божьей Матери («Знамение») (1295);
празднование Абалакской иконы Божией Матери («Знамение») (1637);
празднование Царскосельской иконы Божией Матери («Знамение»);
празднование Серафимо-Понетаевской иконы Божией Матери («Знамение») (1879);
празднование Верхнетагильской иконы Божией Матери («Знамение») (1753);
празднование Корчемной иконы Божией Матери («Знамение») (XVIII в.).


=== Именины ===
Православные: Всеволод, Роман, Яков.

== Приметы ==
Романов день. Роман Чудотворец.
На Романа рыбы ложатся в свои зимовальные ямы, на дно.
По тучам да по звездам гадают о будущей погоде.
Если на заре лицом к северному ветру встать, то сметет он с тебя все надсады, все тяготы.
В это время у лосей отпадают старые рога, а в берлоге засыпает медведь.
`
	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Международные_
- День прав человека
- Всемирный день футбола

_Национальные_
- Россия
- Марий Эл — День марийской письменности
- Таиланд — День Конституции

_Религиозные_
- празднование иконы Божией Матери, именуемой «Знамение» (правосл.)
- обретение мощей благоверного князя Всеволода, во Святом Крещении Гавриила, Новгородского, Псковского чудотворца (1192) (правосл.)
- Собор новомучеников и исповедников Радонежских (правосл.)
- празднование Курской-Коренной иконы Божьей Матери («Знамение») (1295) (правосл.)
- празднование Абалакской иконы Божией Матери («Знамение») (1637) (правосл.)
- празднование Царскосельской иконы Божией Матери («Знамение») (правосл.)
- празднование Серафимо-Понетаевской иконы Божией Матери («Знамение») (1879) (правосл.)
- празднование Верхнетагильской иконы Божией Матери («Знамение») (1753) (правосл.)
- празднование Корчемной иконы Божией Матери («Знамение») (XVIII в.) (правосл.)

_Именины_
- Православные: Всеволод, Роман, Яков

*Приметы*

_Романов день_
Роман Чудотворец
На Романа рыбы ложатся в свои зимовальные ямы, на дно
По тучам да по звездам гадают о будущей погоде
Если на заре лицом к северному ветру встать, то сметет он с тебя все надсады, все тяготы
В это время у лосей отпадают старые рога, а в берлоге засыпает медведь
`

	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_13_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 13, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `13 декабря — 347-й день года (348-й в високосные годы) в григорианском календаре.
До конца года остаётся 18 дней.
В XX и XXI веках соответствует 30 ноября юлианского календаря.


== Праздники и памятные дни ==


=== Национальные ===
 Мальта — День республики.


=== Религиозные ===
 Православиепамять апостола Андрея Первозванного (ок. 62);
память святителя Фрументия, архиепископа Индийского (Эфиопского) (ок. 380);
память священномученика Иоанна Честнова, пресвитера (1937).
 Другие конфессииДень Святой Лючии
День святых Фаддея и Варфоломея в Армении.


=== Именины ===
Православные: Андрей.


== Приметы ==
Андреева ночь. Гадальный день. Святой Андрей. Андрей Первозванный.

В старину на Андрея наслушивали воду: поутру шли на реку, рубили прорубь, прежде чем зачерпнуть воды, опускались на колени на краю проруби, прижимались ухом ко льду и слушали:Когда шумная вода, то надо ждать метели, стужи.
Когда тихая вода на Андреев день, то зима будет тихой, хорошей.


`
	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Национальные_
- Мальта — День республики

_Религиозные_
- память апостола Андрея Первозванного (ок. 62) (правосл.)
- День Святой Лючии
- День святых Фаддея и Варфоломея в Армении

_Именины_
- Православные: Андрей

*Приметы*

_Андреева ночь_
Гадальный день
Святой Андрей
Андрей Первозванный
В старину на Андрея наслушивали воду: поутру шли на реку, рубили прорубь, прежде чем зачерпнуть воды, опускались на колени на краю проруби, прижимались ухом ко льду и слушали:Когда шумная вода, то надо ждать метели, стужи
Когда тихая вода на Андреев день, то зима будет тихой, хорошей
`
	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_18_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 18, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `18 декабря — 352-й день года (353-й в високосные годы) в григорианском календаре.
До конца года остаётся 13 дней.
В XX и XXI веках соответствует 5 декабря юлианского календаря.


== Праздники и памятные дни ==
См. также: Категория:Праздники 18 декабря


=== Международные ===
 Международный день мигранта.
 Международный день арабского языка.


=== Национальные ===
 Катар — День объединения исламских учебных учреждений и университетов в Иране.
 Катар — Национальный день.
 Нигер — День республики.


=== Профессиональные ===
 Молдавия — День полиции.
 Россия — День работников органов ЗАГСа.
 Россия — День подразделений собственной безопасности органов внутренних дел Российской Федерации.


=== Религиозные ===
 Православиепамять преподобного Саввы Освященного (532);
память святителя Гурия, архиепископа Казанского (1563);
память мученика Анастасия Аквилейского;
память преподобных Кариона монаха и сына его Захарии, египтян (IV в.);
память священномученика Илии Четверухина, пресвитера (1932);
память преподобномученика Геннадия (Летюка), иеромонаха (1941);
память священноисповедника Сергия Правдолюбова, пресвитера (1950).


=== Именины ===
Православные: Анастасий, Гурий, Захар, Карион, Нектарий, Савва.


== События ==`

	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Международные_
- Международный день мигранта
- Международный день арабского языка

_Национальные_
- Катар — День объединения исламских учебных учреждений и университетов в Иране
- Катар — Национальный день
- Нигер — День республики

_Профессиональные_
- Молдавия — День полиции
- Россия — День работников органов ЗАГСа
- Россия — День подразделений собственной безопасности органов внутренних дел Российской Федерации

_Именины_
- Православные: Анастасий, Гурий, Захар, Карион, Нектарий, Савва
`
	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_19_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 19, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `19 декабря — 353-й день года (354-й в високосные годы) в григорианском календаре.
До конца года остаётся 12 дней.
В XX и XXI веках соответствует 6 декабря юлианского календаря.


== Праздники и памятные дни ==


=== Профессиональные ===
 Россия — День подразделений военной контрразведки Федеральной Службы Безопасности Российской Федерации.
 Украина — День адвокатуры.


=== Религиозные ===
 Православие
— память святителя Николая, архиепископа Мир Ликийских, чудотворца (ум. ок. 335)
— память cвятителя Феофила исповедника, епископа Антиохийского (ум. 181)
— память святителя Николая, епископа Патарского (ум. IV)
— блаженного Максима, митрополита Киевского и всея Руси (ум. 1305)
— память мученика Николая Карамана, Смирнскаго (ум. 1657)


== События ==

`

	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Профессиональные_
- Россия — День подразделений военной контрразведки Федеральной Службы Безопасности Российской Федерации
- Украина — День адвокатуры

_Религиозные_
- блаженного Максима, митрополита Киевского и всея Руси (ум. 1305) (правосл.)
`
	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}

func TestParse_24_12(t *testing.T) {
	//location, _ := time.LoadLocation("Europe/Moscow")
	//now := time.Date(2018, time.December, 24, 1, 1, 1, 1, location)
	//fullReport := getWikiReport(&now)

	fullReport := `24 декабря — 358-й день года (359-й в високосные годы) в григорианском календаре.
До конца года остаётся 7 дней.
В XX и XXI веках соответствует 11 декабря юлианского календаря.


== Праздники и памятные дни ==


=== Национальные ===
 Россия
День воинской славы России — День взятия турецкой крепости Измаил российскими войсками под командованием А. В. Суворова (1790).
День ФАПСИ.
 Украина — День работников архивных учреждений.
 Приднестровье — День Конституции.


=== Религиозные ===
 Рождественский сочельник у католиков и протестантов, отмечающих Рождество по григорианскому календарю, и у православных, отмечающих Рождество по новоюлианскому календарю.

 ПравославиеНеделя святых праотец (переходящее празднование в 2017 году);
память преподобного Даниила Столпника (493);
память преподобного Никона Сухого, Печерского, в Ближних пещерах (XII в.);
память мученика Миракса Египтянина (VII);
память мучеников Акепсия и Аифала;
память преподобного Луки Столпника, иеромонаха (ок. 970—980);
память священномученика Феофана (Ильминского), епископа Соликамского, и с ним двух священномучеников и пяти мучеников (1918);
память священномученика Николая Виноградова, пресвитера (1937);
память священномученика Иоанна Богоявленского, пресвитера (1941).


=== Именины ===
Адам, Александр, Никон, Даниил, Емельян, Иван, Леонтий, Николай, Пётр, Терентий

`
	report2, _ := Parse(fullReport)
	expected := `*Праздники и памятные дни*

_Национальные_
- Россия
- День воинской славы России — День взятия турецкой крепости Измаил российскими войсками под командованием А. В. Суворова (1790)
- День ФАПСИ
- Украина — День работников архивных учреждений
- Приднестровье — День Конституции

_Религиозные_
- Неделя святых праотец (переходящее празднование в 2017 году) (правосл.)
- Рождественский сочельник у католиков и протестантов, отмечающих Рождество по григорианскому календарю, и у православных, отмечающих Рождество по новоюлианскому календарю

_Именины_
- Адам, Александр, Никон, Даниил, Емельян, Иван, Леонтий, Николай, Пётр, Терентий
`
	s2 := report2.String()
	if expected != s2 {
		t.Error(expected, "\n!=\n", s2)
	}
}
