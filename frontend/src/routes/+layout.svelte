<!-- src/routes/+layout.svelte -->
<script lang="ts">
	import './layout.css';
	import Navbar from '$lib/components/Navbar.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import { authState } from '$lib/state/auth.svelte';
	import { onMount } from 'svelte';

	let { children } = $props();

	onMount(() => {
		authState.checkAuth();
	});
</script>

<svelte:head>
	<title>UniApp - University Admission Portal</title>
</svelte:head>

<div class="min-h-screen flex flex-col bg-slate-50 text-slate-900 font-sans antialiased">
	<Navbar />
	<main class="flex-1">
		{@render children()}
	</main>
	<Footer />
</div>
