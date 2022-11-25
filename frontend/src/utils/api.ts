import Axios from 'axios'
const token = new URLSearchParams(window.location.search).get('token')
export const axios = Axios.create({
	baseURL: 'http://ictc-int.sit.kmutt.ac.th:3000/api',
	// baseURL: 'http://10.0.222.181:3000/api',
	headers: {
		common: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/json',
		},
	},
})
