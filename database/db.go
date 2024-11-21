package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS ei_jobs")
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	_, err = db.Exec("USE ei_jobs")
	if err != nil {
		return fmt.Errorf("failed to switch database: %w", err)
	}

	tableQueries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			first_name VARCHAR(50) NULL,
			last_name VARCHAR(50) NULL,
			company_name VARCHAR(50) NULL,
			avatar_url VARCHAR(50) NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			phone VARCHAR(100) NOT NULL UNIQUE,
			role_id INT,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS roles (
		    id INT AUTO_INCREMENT PRIMARY KEY,
		    name VARCHAR(50) NOT NULL UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS resumes (
            id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT,
            date_of_birth DATE NOT NULL,
            gender VARCHAR(100) NOT NULL,
            specialization_id INT NOT NULL,
            description VARCHAR(10000) NULL,
            salary_from INT NULL,
            salary_to INT NULL,
            salary_period VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
		`CREATE TABLE IF NOT EXISTS resume_skills (
            id INT AUTO_INCREMENT PRIMARY KEY,
            resume_id INT,
            skill VARCHAR(255)
        )`,
		`CREATE TABLE IF NOT EXISTS resume_organizations (
            id INT AUTO_INCREMENT PRIMARY KEY,
            resume_id INT NOT NULL,
            oraganization_name VARCHAR(255) NOT NULL,
            specialization_id INT NOT NULL,
            description VARCHAR(1000) NOT NULL,
            start_month VARCHAR(255) NOT NULL,
            start_year VARCHAR(255) NOT NULL,
            end_month VARCHAR(255) NULL,
            end_year VARCHAR(255) NULL
        )`,
		`CREATE TABLE IF NOT EXISTS vacancies (
            id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL,
            specialization_id INT NOT NULL,
            title VARCHAR(255) NOT NULL,
            country VARCHAR(255) NOT NULL,
            city VARCHAR(255) NOT NULL,
            salary_from INT NULL,
            salary_to INT NULL,
            salary_period VARCHAR(255) NOT NULL,
            work_format VARCHAR(255) NOT NULL,
            work_schedule VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`,
		`CREATE TABLE IF NOT EXISTS vacancy_conditions (
        	id INT AUTO_INCREMENT PRIMARY KEY,
         	vacancy_id INT NOT NULL,
          	icon VARCHAR(255) NOT NULL,
           	condition_text VARCHAR(1000) NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS vacancy_requirements (
           	id INT AUTO_INCREMENT PRIMARY KEY,
           	vacancy_id INT NOT NULL,
           	requirement VARCHAR(1000) NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS vacancy_responsibilities (
           	id INT AUTO_INCREMENT PRIMARY KEY,
           	vacancy_id INT NOT NULL,
           	responsibility VARCHAR(1000) NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS specialization_categories (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE
        )`,
		`CREATE TABLE IF NOT EXISTS specializations (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            category_id INT NOT NULL,
            FOREIGN KEY (category_id) REFERENCES specialization_categories(id)
        )`,
		`CREATE TABLE IF NOT EXISTS chats (
        	id INT AUTO_INCREMENT PRIMARY KEY,
         	user_id INT NOT NULL,
          	receiver_id INT NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS user_services (
        	id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL,
            name VARCHAR(255) NOT NULL,
            description VARCHAR(1000) NULL,
            price INT NOT NULL,
            deadline VARCHAR(255) NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS user_description (
        	id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL,
            description VARCHAR(10000) NOT NULL
        )`,
	}

	for _, query := range tableQueries {
		_, err = db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	log.Println("Database migration completed successfully")
	return nil
}

func SeedDatabase(db *sql.DB) error {
	specializationCategoryQueries := []string{
		`INSERT INTO specialization_categories (name) VALUES
            ('Информационные технологии'),
            ('Здравоохранение'),
            ('Инженерия'),
            ('Финансы'),
            ('Маркетинг'),
            ('Управление персоналом'),
            ('Образование'),
            ('Юриспруденция'),
            ('Искусство и дизайн'),
            ('Гостеприимство'),
            ('Транспорт'),
            ('Строительство'),
            ('Сельское хозяйство'),
            ('Производство'),
            ('Розничная торговля'),
            ('Телекоммуникации'),
            ('Энергетика'),
            ('Экология'),
            ('Государственное управление'),
            ('Некоммерческие организации')
        `,
	}

	for _, query := range specializationCategoryQueries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to insert specialization categories: %w", err)
		}
	}

	specializationQueries := []string{
		`INSERT INTO specializations (name, category_id) VALUES
            -- Информационные технологии
            ('Инженер-программист', 1),
            ('Веб-разработчик', 1),
            ('Администратор баз данных', 1),
            ('Сетевой администратор', 1),
            ('Аналитик по кибербезопасности', 1),
            ('Датасайентист', 1),
            ('IT-менеджер проектов', 1),
            ('Аналитик компьютерных систем', 1),
            ('Программист', 1),
            ('Специалист технической поддержки', 1),

            -- Здравоохранение
            ('Медицинская сестра', 2),
            ('Врач', 2),
            ('Фармацевт', 2),
            ('Физиотерапевт', 2),
            ('Зубной гигиенист', 2),
            ('Лаборант', 2),
            ('Помощник медсестры', 2),
            ('Рентгенолог', 2),
            ('Эрготерапевт', 2),
            ('Специалист по медицинской документации', 2),

            -- Инженерия
            ('Инженер-механик', 3),
            ('Инженер-строитель', 3),
            ('Инженер-электрик', 3),
            ('Инженер-конструктор', 3),
            ('Инженер-эколог', 3),
            ('Инженер-технолог', 3),
            ('Авиационный инженер', 3),
            ('Химический инженер', 3),
            ('Биомедицинский инженер', 3),
            ('Инженер-проектировщик', 3),

            -- Финансы
            ('Бухгалтер', 4),
            ('Финансовый аналитик', 4),
            ('Финансовый менеджер', 4),
            ('Инвестиционный банкир', 4),
            ('Финансовый консультант', 4),
            ('Аудитор', 4),
            ('Аналитик бюджета', 4),
            ('Налоговый специалист', 4),
            ('Кредитный аналитик', 4),
            ('Страховой андеррайтер', 4),

            -- Маркетинг
            ('Маркетинг-менеджер', 5),
            ('Специалист по социальным медиа', 5),
            ('Стратег цифрового маркетинга', 5),
            ('Контент-писатель', 5),
            ('Оптимизатор поисковых систем', 5),
            ('Бренд-менеджер', 5),
            ('Организатор мероприятий', 5),
            ('Аналитик рыночных исследований', 5),
            ('Специалист по связям с общественностью', 5),
            ('Графический дизайнер', 5),

            -- Управление персоналом
            ('Менеджер по персоналу', 6),
            ('Рекрутер', 6),
            ('Специалист по обучению и развитию', 6),
            ('Специалист по взаимоотношениям с сотрудниками', 6),
            ('Аналитик компенсаций и льгот', 6),
            ('Консультант по организационному развитию', 6),
            ('Специалист по привлечению талантов', 6),
            ('Аналитик HRIS', 6),
            ('HR-генералист', 6),
            ('Координатор по разнообразию и инклюзии', 6),

            -- Образование
            ('Учитель начальной школы', 7),
            ('Учитель старшей школы', 7),
            ('Профессор', 7),
            ('Дизайнер обучения', 7),
            ('Школьный консультант', 7),
            ('Учитель специального образования', 7),
            ('Воспитатель детского сада', 7),
            ('Библиотекарь', 7),
            ('Администратор образовательного учреждения', 7),
            ('Репетитор', 7),

            -- Юриспруденция
            ('Юрист', 8),
            ('Парапрофессионал', 8),
            ('Помощник юриста', 8),
            ('Специалист по договорам', 8),
            ('Специалист по соблюдению требований', 8),
            ('Юридический исследователь', 8),
            ('Патентный агент', 8),
            ('Специалист по судебному делопроизводству', 8),
            ('Юридический секретарь', 8),
            ('Специалист по вопросам иммиграции', 8),

            -- Искусство и дизайн
            ('Графический дизайнер', 9),
            ('Дизайнер интерьера', 9),
            ('Модельер', 9),
            ('Иллюстратор', 9),
            ('UX/UI-дизайнер', 9),
            ('Фотограф', 9),
            ('Арт-директор', 9),
            ('Веб-дизайнер', 9),
            ('Аниматор', 9),
            ('Копирайтер', 9),

            -- Гостеприимство
            ('Менеджер отеля', 10),
            ('Повар', 10),
            ('Менеджер ресторана', 10),
            ('Организатор мероприятий', 10),
            ('Туристический агент', 10),
            ('Консьерж', 10),
            ('Бармен', 10),
            ('Менеджер кейтеринга', 10),
            ('Экскурсовод', 10),
            ('Спа-терапевт', 10),

            -- Транспорт
            ('Авиационный пилот', 11),
            ('Авиадиспетчер', 11),
            ('Менеджер логистики', 11),
            ('Дальнобойщик', 11),
            ('Капитан корабля', 11),
            ('Машинист поезда', 11),
            ('Курьер', 11),
            ('Специалист по транспортному планированию', 11),
            ('Автомеханик', 11),
            ('Менеджер склада', 11),

            -- Строительство
            ('Менеджер строительства', 12),
            ('Архитектор', 12),
            ('Электрик', 12),
            ('Сантехник', 12),
            ('Плотник', 12),
            ('Сварщик', 12),
            ('Специалист по HVAC', 12),
            ('Ландшафтный архитектор', 12),
            ('Прораб', 12),
            ('Строительный инспектор', 12),

            -- Сельское хозяйство
            ('Фермер', 13),
            ('Сельскохозяйственный ученый', 13),
            ('Агроном', 13),
            ('Менеджер фермы', 13),
            ('Садовод', 13),
            ('Оператор сельхозтехники', 13),
            ('Пищевой технолог', 13),
            ('Ветеринар', 13),
            ('Специалист по борьбе с вредителями', 13),
            ('Специалист по ирригации', 13),

            -- Производство
            ('Менеджер производства', 14),
            ('Инженер-технолог', 14),
            ('Специалист по контролю качества', 14),
            ('Станочник', 14),
            ('Техник-робототехник', 14),
            ('Аналитик цепочки поставок', 14),
            ('Техник-оператор', 14),
            ('Специалист по улучшению процессов', 14),
            ('Материаловед', 14),
            ('Техник-механик', 14),

            -- Розничная торговля
            ('Менеджер розничной торговли', 15),
            ('Продавец-консультант', 15),
            ('Мерчендайзер', 15),
            ('Аналитик инвентаризации', 15),
            ('Визуальный мерчендайзер', 15),
            ('Специалист по работе с клиентами', 15),
            ('Закупщик', 15),
            ('Специалист по предотвращению потерь', 15),
            ('Специалист по электронной коммерции', 15),
            ('Специалист по планировке магазина', 15),

            -- Телекоммуникации
            ('Сетевой инженер', 16),
            ('Техник в области телекоммуникаций', 16),
            ('Менеджер проекта в телекоммуникациях', 16),
            ('Специалист по беспроводным сетям', 16),
            ('Аналитик в области телекоммуникаций', 16),
            ('Техник широкополосных сетей', 16),
            ('Торговый представитель в телекоммуникациях', 16),
            ('Техник оптоволоконных сетей', 16),
            ('Специалист по VoIP', 16),
            ('Консультант в области телекоммуникаций', 16),

            -- Энергетика
            ('Оператор электростанции', 17),
            ('Аудитор по энергоэффективности', 17),
            ('Специалист по возобновляемым источникам энергии', 17),
            ('Инженер нефтегазовой отрасли', 17),
            ('Менеджер коммунального хозяйства', 17),
            ('Консультант по энергоэффективности', 17),
            ('Техник геотермальной энергетики', 17),
            ('Распределитель электроэнергии', 17),
            ('Техник ветровых турбин', 17),
            ('Установщик солнечных панелей', 17),

            -- Экология
            ('Эколог', 18),
            ('Координатор по устойчивому развитию', 18),
            ('Инженер-эколог', 18),
            ('Лесной инспектор', 18),
            ('Специалист по управлению отходами', 18),
            ('Эковолонтер', 18),
            ('Эколог', 18),
            ('Техник по контролю загрязнения', 18),
            ('Специалист по экологическому соответствию', 18),
            ('Экологический консультант', 18),

            -- Государственное управление
            ('Государственный администратор', 19),
            ('Аналитик политики', 19),
            ('Городской планировщик', 19),
            ('Социальный работник', 19),
            ('Специалист по общественному здравоохранению', 19),
            ('Криминолог', 19),
            ('Городской планировщик', 19),
            ('Аналитик государственных финансов', 19),
            ('Специалист по управлению чрезвычайными ситуациями', 19),
            ('Политолог', 19),

            -- Некоммерческие организации
            ('Fundraiser', 20),
            ('Грант-райтер', 20),
            ('Менеджер программ НКО', 20),
            ('Координатор волонтеров', 20),
            ('Специалист по защите прав', 20),
            ('Исполнительный директор НКО', 20),
            ('Бухгалтер НКО', 20),
            ('Специалист по персоналу в НКО', 20),
            ('Маркетолог и специалист по связям с общественностью в НКО', 20),
            ('Сотрудник отдела развития НКО', 20)
        `,
	}

	for _, query := range specializationQueries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to insert specializations: %w", err)
		}
	}

	return nil
}
