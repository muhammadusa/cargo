const ordersElement = document.getElementById("ordersElement")

async function main () {

	const response = await fetch('http://localhost:8080')

	const orders = await response.json()

	for (const order of orders) {

		const liElement = document.createElement("li")

		ordersElement.appendChild(liElement)

		const pElement = document.createElement("p")

		liElement.appendChild(pElement)

		const userElement = document.createElement("span")
		userElement.textContent = order.Name

		const phoneElement = document.createElement("span")
		phoneElement.textContent = order.Phone

		pElement.appendChild(userElement)
		pElement.appendChild(phoneElement)
	}
}

nameInput.value = null
phoneInput.value = null

main()

sendBtn.onclick = async function () {

	const response = await fetch('http://localhost:4000/orders', {
		method: 'POST',
		body: JSON.stringify({ Name: nameInput.value, Phone: phoneInput.value, }),
	})

	if (response.status === 201) {

		console.log('Created!', response)
	}

	const neworder = await response.json()

	const newLi = document.createElement('LI')
	const p = document.createElement('p')
	const span1 = document.createElement('span')
	const span2 = document.createElement('span')

	p.appendChild(span1)
	p.appendChild(span2)

	newLi.appendChild(p)

	ordersElement.appendChild(newLi)

	span1.textContent = neworder.Name
	span2.textContent = neworder.Phone

	nameInput.value = null
	phoneInput.value = null
}