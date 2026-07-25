// src/lib/api/landing.ts
// Demo data — replace fetch() calls with real backend URLs when ready.

import type { Stat, Feature, Testimonial } from '$lib/types/landing';

// ---------------------------------------------------------------------------
// Stats
// ---------------------------------------------------------------------------

const DEMO_STATS: Stat[] = [
	{ value: '50+', label: 'Partner Universities' },
	{ value: '10k', label: 'Active Students' },
	{ value: '95%', label: 'Match Accuracy' },
	{ value: '24h', label: 'Support Response' }
];

export async function fetchStats(): Promise<Stat[]> {
	// TODO: replace with real endpoint
	// const res = await fetch('/api/stats');
	// return res.json();
	await new Promise((r) => setTimeout(r, 300)); // simulate latency
	return DEMO_STATS;
}

// ---------------------------------------------------------------------------
// Features
// ---------------------------------------------------------------------------

const DEMO_FEATURES: Feature[] = [
	{
		icon: 'troubleshoot',
		iconBg: 'bg-primary-fixed',
		iconColor: 'text-primary',
		title: 'Smart Eligibility Engine',
		description:
			'Instantly see which universities you qualify for based on your academic profile. Stop guessing and start applying with confidence.',
		cta: 'Try Engine',
		wide: true
	},
	{
		icon: 'route',
		iconBg: 'bg-tertiary-fixed',
		iconColor: 'text-tertiary',
		title: 'Route Tracker',
		description: 'Visualize your application journey from submission to enrollment.'
	},
	{
		icon: 'notifications_active',
		iconBg: 'bg-secondary-fixed',
		iconColor: 'text-secondary',
		title: 'Live Alerts',
		description: 'Never miss a deadline or exam date with synchronized notifications.'
	},
	{
		icon: 'payments',
		iconBg: 'bg-error-container',
		iconColor: 'text-error',
		title: 'Unified Payments',
		description:
			'Handle all application fees and document verification costs through a single, secure gateway. Track your spending effortlessly.',
		wide: true
	}
];

export async function fetchFeatures(): Promise<Feature[]> {
	// TODO: replace with real endpoint
	// const res = await fetch('/api/features');
	// return res.json();
	await new Promise((r) => setTimeout(r, 200));
	return DEMO_FEATURES;
}

// ---------------------------------------------------------------------------
// Testimonials
// ---------------------------------------------------------------------------

const DEMO_TESTIMONIALS: Testimonial[] = [
	{
		quote:
			'UniApp made finding eligible universities incredibly easy. The interface is clean, and the recommendations were spot on. Highly recommended for stressed students!',
		name: 'Rahim Uddin',
		acceptance: 'Accepted to BUET',
		avatarUrl:
			'https://lh3.googleusercontent.com/aida-public/AB6AXuBLziJYhQnn30-MTMtOT99AVJamee_uzadOU7t6j80_pzscLftLLFaoFjS9BfZC2FqhBpsF5yvJvQU_AwIfCLSH_ciDda_1Gqo_fU_83gTiPXhO9el-ORu5XiE3XFqEvvoFdgBA0f7ckB776EHUjF7rubCKOKkyRykGJSwbjpprWLTB-50r6PvZvgd864LkH3LlvLscEVeDusFxBL38nAHowB1H9TDK70TyGk_OJt2hegiq4DA_1QMUPNI6J3pv7fdnymxSp2pxVy7-'
	},
	{
		quote:
			'The unified application feature saved me weeks of repetitive typing. I could focus on my exams instead of filling out the same forms twenty times.',
		name: 'Sarah Ahmed',
		acceptance: 'Accepted to DU',
		avatarUrl:
			'https://lh3.googleusercontent.com/aida-public/AB6AXuDR2xT4ceh2YdfWOjgnRQspgxJG-u2E8UI3sdVwP8FVJRZIHwSSTeLiZYyXZ_oIT7CKwdd8enfbud-NqPGTtWxkwfeBzCZySEhfsBA3-3ba2JvHhjLMG3GgA26WsFeqxgwQv4tEAZXla-5qeU7iUvaFRCTzSEoBaMZ22fRUIYwolud3l072x17M8-vyAWOtwbQjhWuIDoAV4cYIhKN_PZvNPnstNDOa6dlujLOXdGTWL26QAHQhMmfyhi0GILdK1baCz-wQZGu5Uvw3'
	},
	{
		quote:
			'Tracking deadlines across different varsity routes used to be a nightmare. The calendar and alerts on this platform are lifesavers.',
		name: 'Fahad Hossain',
		acceptance: 'Accepted to RUET',
		avatarUrl:
			'https://lh3.googleusercontent.com/aida-public/AB6AXuDM0VVrAXoa4v8Jpt2tKvJcAzl06t1LdrTrwx2LTIcI26wPw6k4jPmylqOGgwYgBvvWQ8rQDwvQJyITeMLys4P3taLjcFPsS7cganwzJBL3xE-iFYcImPJo8_P182TFN3QtS-46bgloNoCNUPzyTZRCwoEGkjiomPjB4Ok8XFoiwRQ3vk1lIGVgDlI6fUCIpOOnlu-7fb5r3tOOaVxDtYERWm-ACEIySWYkn71PFGNdYJPt4vmgTT3Ouz9BRAyrvQDQHo01FlWwcuRB'
	}
];

export async function fetchTestimonials(): Promise<Testimonial[]> {
	// TODO: replace with real endpoint
	// const res = await fetch('/api/testimonials');
	// return res.json();
	await new Promise((r) => setTimeout(r, 250));
	return DEMO_TESTIMONIALS;
}
