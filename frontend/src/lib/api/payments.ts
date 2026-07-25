// src/lib/api/payments.ts

import { apiFetch } from './client';
import type { PaymentRequest, PaymentResponse } from '$lib/types/models';

export async function processPayment(payload: PaymentRequest): Promise<PaymentResponse> {
	return apiFetch<PaymentResponse>('/payments/process', {
		method: 'POST',
		body: JSON.stringify(payload)
	});
}
