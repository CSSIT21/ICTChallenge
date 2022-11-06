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
			keyframes: {
				float: {
					'0%, 100%': {
						transform: 'translateY(0)',
					},
					'50%': {
						transform: 'translateY(-30px)',
					},
				},
			},
			animation: {
				floating: 'float 6s ease-in-out infinite',
			},
		},
	},

	plugins: [],
}

module.exports = config
