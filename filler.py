import random
import requests
from faker import Faker

fake = Faker('ru_RU')

url = "http://localhost:5000/api/"
headers = {'content-type': 'application/json'}
count = 1000
genders = ['М', 'Ж', 'Н']
haircuts = {
	'Каре' : 'женская, короткая, простая',
	'Каскад' : 'женская, короткая, ассиметричность, хаотичность',
	'Боб' : 'женская, короткая, мягкие контуры',
	'Шапочка' : 'женская, короткая, актуальная, эффектная',
	'Каскад длинный' : 'женская, длинная, градуированные пряди',
	'Косичка' : 'женская, длинная, косичка',
	'Как у актрисы' : 'женская, актриса на выбор',
	'Налысо' : 'бывает и такое',
	'Собрался на военку' : 'мужская, ультракороткая, Кабардинский одобряет',
	'Как у актера' : 'мужская, актер на выбор',
	'По фотографии' : 'зависит от вашего выбора'
}

roles = ['junior', 'middle', 'senior', 'admin']

def create_salons():
	try:
		response = requests.post(url + "salon/create", json=
		{
			"address": fake.address()
		})
		print(response)
	except Exception as e:
		print(e)

def create_customer():
	try:
		response = requests.post(url + "customer/create", json=
		{
			"fullname": fake.name(),
			"email": fake.email(),
			"gender": random.choice(genders),
			"address": fake.address()
		})
		print(response)
	except Exception as e:
		print(e)

def create_haircut(i):
	try:
		name = list(haircuts.keys())[i]
		description = haircuts[name]
		response = requests.post(url + "haircut/create", json=
		{
			"name": name,
			"description": description
		})
		print(response)
	except Exception as e:
		print(e)

def create_employee():
	try:
		response = requests.post(url + "employee/create", json=
		{
			"fullname": fake.name(),
			"role": random.choice(roles),
			"status": fake.text(max_nb_chars=200, ext_word_list=None),
			"email": fake.email(),
			"gender": random.choice(genders),
			"is_working": fake.pybool(),
			"salon_id": random.randint(1, count + 1)
		})
		print(response)
	except Exception as e:
		print(e)

def create_price(i):
	try:
		response = requests.post(url + "price/create", json=
		{
			"value" : random.randint(500, 10000),
			"haircut_id": i,
			"date": fake.unix_time(end_datetime=None, start_datetime=None)
		})
		print(response)
	except Exception as e:
		print(e)

def create_deals(i):
	try:
		response = requests.post(url + "deal/create", json=
		{
			"customer_id" : random.randint(500, 10000),
			"haircut_id": random.randint(1, 11),
			"employee_id": random.randint(1, 1001),
			"price_id": random.randint(1, 11)
		})
		print(response)
	except Exception as e:
		print(e)

# for _ in range(10 * count):
# 	create_customer()
#
# for _ in range(count):
# 	create_salons()
#
# for k in range(len(haircuts)):
# 	create_haircut(k)
#
# for _ in range(3 * count):
# 	create_employee()
# 
# for k in range(len(haircuts)):
# 	create_price(k)

for k in range(len(haircuts)):
	create_deals(k)
