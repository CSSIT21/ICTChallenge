const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		boxShadow: {
			DEFAULT: '8px 10px 30px rgba(135, 135, 135, 0.25);',
		},
		extend: {
			backgroundImage: {
				'preview-bg': "url('/src/assets/images/preview-bg.svg')",
			},
		},
	},

	plugins: [],
}

module.exports = config
