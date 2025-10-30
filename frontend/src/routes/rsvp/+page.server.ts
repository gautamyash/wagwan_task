import type { Actions } from './$types';
import { fail } from '@sveltejs/kit';

export const actions: Actions = {
  default: async ({ request, fetch }) => {
    const formData = await request.formData();

    const name = String(formData.get('name') || '').trim();
    const email = String(formData.get('email') || '').trim();
    const phone = String(formData.get('phone') || '').trim();
    const status = String(formData.get('status') || 'attending');
    const notes = String(formData.get('notes') || '').trim();
    const plus_ones = Number(formData.get('plus_ones') || 0);
    const dietary_restrictions = String(formData.get('dietary_restrictions') || '').trim();

    // Basic validation
    if (!name || !email || !phone) {
      return fail(400, {
        error: 'Please fill in all required fields.',
        values: { name, email, phone, status, notes, plus_ones, dietary_restrictions }
      });
    }
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      return fail(400, {
        error: 'Please enter a valid email address.',
        values: { name, email, phone, status, notes, plus_ones, dietary_restrictions }
      });
    }

    // Proxy to Go API
    const response = await fetch('http://localhost:8080/api/guests', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name,
        email,
        phone,
        status,
        notes,
        plus_ones,
        dietary_restrictions,
        event_id: 1
      })
    });

    if (!response.ok) {
      let message = 'Submission failed';
      try {
        const err = await response.json();
        message = err?.error || message;
      } catch (_) {}
      return fail(500, { error: message });
    }

    return {
      success: true,
      message: '🎉 Thank you! Your RSVP has been recorded.'
    };
  }
};


