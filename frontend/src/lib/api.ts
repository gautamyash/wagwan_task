const API_BASE_URL = 'http://localhost:8080/api';

export interface Guest {
	id: number;
	name: string;
	email: string;
	phone: string;
	status: 'pending' | 'attending' | 'declined';
	notes?: string;
	plus_ones?: number;
	dietary_restrictions?: string;
	rsvp_date?: string;
	created_at: string;
}

export interface CreateGuestData {
	name: string;
	email: string;
	phone: string;
	status?: 'pending' | 'attending' | 'declined';
}

export async function getGuests(statusFilter?: string): Promise<Guest[]> {
	let url = `${API_BASE_URL}/guests`;
	
	if (statusFilter) {
		url += `?status=${statusFilter}`;  // Fixed: Changed 'filter' to 'status' to match backend
	}

	const response = await fetch(url);
	
	if (!response.ok) {
		throw new Error('Failed to fetch guests');
	}

	return response.json();
}

export async function getGuest(id: number): Promise<Guest> {
	const response = await fetch(`${API_BASE_URL}/guests/${id}`);
	
	if (!response.ok) {
		throw new Error('Failed to fetch guest');
	}

	return response.json();
}

export async function createGuest(data: CreateGuestData): Promise<Guest> {
	const response = await fetch(`${API_BASE_URL}/guests`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(data)
	});

	if (!response.ok) {
		const error = await response.json();
		throw new Error(error.error || 'Failed to create guest');
	}

	return response.json();
}

export async function deleteGuest(id: number): Promise<void> {
	const response = await fetch(`${API_BASE_URL}/guests/${id}`, {
		method: 'DELETE'
	});

	if (!response.ok) {
		throw new Error('Failed to delete guest');
	}
}

