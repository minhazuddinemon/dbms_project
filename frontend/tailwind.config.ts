import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	darkMode: 'class',
	theme: {
		extend: {
			colors: {
				outline: '#767586',
				'surface-dim': '#d8dadc',
				'surface-variant': '#e0e3e5',
				tertiary: '#006c49',
				'secondary-fixed-dim': '#c3c0ff',
				'on-primary': '#ffffff',
				'surface-tint': '#494bd6',
				'primary-fixed': '#e1e0ff',
				'on-background': '#191c1e',
				error: '#ba1a1a',
				'inverse-primary': '#c0c1ff',
				'on-tertiary-fixed': '#002113',
				primary: '#4648d4',
				'error-container': '#ffdad6',
				'on-tertiary': '#ffffff',
				'on-secondary-fixed-variant': '#372abf',
				'tertiary-fixed': '#6ffbbe',
				'inverse-on-surface': '#eff1f3',
				'on-tertiary-fixed-variant': '#005236',
				'surface-container-low': '#f2f4f6',
				'primary-container': '#6063ee',
				'on-primary-fixed-variant': '#2f2ebe',
				'surface-container-lowest': '#ffffff',
				'on-primary-fixed': '#07006c',
				'tertiary-container': '#00885d',
				'on-error-container': '#93000a',
				background: '#f7f9fb',
				'on-secondary-fixed': '#100069',
				'tertiary-fixed-dim': '#4edea3',
				'outline-variant': '#c7c4d7',
				'secondary-fixed': '#e3dfff',
				'inverse-surface': '#2d3133',
				'surface-container-high': '#e6e8ea',
				'on-tertiary-container': '#000703',
				'on-surface': '#191c1e',
				'on-secondary': '#ffffff',
				'on-primary-container': '#fffbff',
				surface: '#f7f9fb',
				'primary-fixed-dim': '#c0c1ff',
				'on-secondary-container': '#fffbff',
				'surface-bright': '#f7f9fb',
				'surface-container-highest': '#e0e3e5',
				'secondary-container': '#6860ef',
				'on-error': '#ffffff',
				'surface-container': '#eceef0',
				secondary: '#4e45d5',
				'on-surface-variant': '#464554'
			},
			borderRadius: {
				DEFAULT: '0.25rem',
				lg: '0.5rem',
				xl: '0.75rem',
				full: '9999px'
			},
			fontFamily: {
				'label-lg': ['Inter', 'sans-serif'],
				'label-sm': ['Inter', 'sans-serif'],
				'headline-md': ['Inter', 'sans-serif'],
				'headline-lg': ['Inter', 'sans-serif'],
				'body-lg': ['Inter', 'sans-serif'],
				'headline-lg-mobile': ['Inter', 'sans-serif'],
				'display-lg': ['Inter', 'sans-serif'],
				'body-md': ['Inter', 'sans-serif']
			},
			fontSize: {
				'label-lg': ['14px', { lineHeight: '20px', letterSpacing: '0.02em', fontWeight: '600' }],
				'label-sm': ['12px', { lineHeight: '16px', letterSpacing: '0.04em', fontWeight: '500' }],
				'headline-md': ['24px', { lineHeight: '32px', fontWeight: '600' }],
				'headline-lg': [
					'32px',
					{ lineHeight: '40px', letterSpacing: '-0.01em', fontWeight: '700' }
				],
				'body-lg': ['18px', { lineHeight: '28px', fontWeight: '400' }],
				'headline-lg-mobile': ['28px', { lineHeight: '36px', fontWeight: '700' }],
				'display-lg': ['48px', { lineHeight: '60px', letterSpacing: '-0.02em', fontWeight: '700' }],
				'body-md': ['16px', { lineHeight: '24px', fontWeight: '400' }]
			},
			animation: {
				'gradient-xy': 'gradient-xy 15s ease infinite',
				float: 'float 15s ease-in-out infinite'
			},
			keyframes: {
				'gradient-xy': {
					'0%, 100%': {
						'background-size': '400% 400%',
						'background-position': 'left center'
					},
					'50%': {
						'background-size': '200% 200%',
						'background-position': 'right center'
					}
				},
				float: {
					'0%, 100%': { transform: 'translateY(0) scale(1)' },
					'50%': { transform: 'translateY(-15px) scale(1.02)' }
				}
			}
		}
	},
	plugins: []
} satisfies Config;
