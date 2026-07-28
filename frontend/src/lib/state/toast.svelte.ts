// src/lib/state/toast.svelte.ts

export interface ToastMessage {
	id: string;
	type: 'success' | 'error' | 'info' | 'warning';
	message: string;
	duration?: number;
}

class ToastState {
	toasts = $state<ToastMessage[]>([]);

	show(type: 'success' | 'error' | 'info' | 'warning', message: string, duration = 4000) {
		const id = Math.random().toString(36).substring(2, 9);
		const toast: ToastMessage = { id, type, message, duration };
		this.toasts = [...this.toasts, toast];

		if (duration > 0) {
			setTimeout(() => {
				this.remove(id);
			}, duration);
		}
	}

	success(message: string, duration = 4000) {
		this.show('success', message, duration);
	}

	error(message: string, duration = 5000) {
		this.show('error', message, duration);
	}

	info(message: string, duration = 4000) {
		this.show('info', message, duration);
	}

	warning(message: string, duration = 4500) {
		this.show('warning', message, duration);
	}

	remove(id: string) {
		this.toasts = this.toasts.filter((t) => t.id !== id);
	}
}

export const toastState = new ToastState();