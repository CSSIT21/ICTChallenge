const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			backgroundImage: {
				'preview-bg': "url('/src/assets/images/preview-bg.svg')",
			},
		},
	},

	plugins: [],
}

module.exports = config
