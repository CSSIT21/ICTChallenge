import { Axios } from 'axios'

export const axios = new Axios({
	baseURL: 'http://localhost:3000/api',
	withCredentials: true,
})
