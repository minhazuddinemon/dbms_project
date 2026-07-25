// src/lib/types/landing.ts

export interface Stat {
	value: string;
	label: string;
}

export interface Feature {
	icon: string;
	iconBg: string;
	iconColor: string;
	title: string;
	description: string;
	cta?: string;
	wide?: boolean; // spans 2 columns on md
}

export interface Testimonial {
	quote: string;
	name: string;
	acceptance: string;
	avatarUrl: string;
}