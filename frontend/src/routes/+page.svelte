<script lang="ts">
	import { onMount } from 'svelte';
	import { getGuests, createGuest, deleteGuest, type Guest } from '$lib/api';

	let guests = $state<Guest[]>([]);
	let loading = $state(true);
	let error = $state<string | null>(null);
	let showForm = $state(false);
	let statusFilter = $state<string>('');

	// Form fields
	let formName = $state('');
	let formEmail = $state('');
	let formPhone = $state('');
	let formStatus = $state<'pending' | 'attending' | 'declined'>('pending');
	let formError = $state<string | null>(null);
	let formSubmitting = $state(false);

	async function loadGuests() {
		try {
			loading = true;
			error = null;
			guests = await getGuests(statusFilter || undefined);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Failed to load guests';
		} finally {
			loading = false;
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		
		// BUG #2: Missing email validation
		// Should validate email format before submission
		if (!formName.trim()) {
			formError = 'Name is required';
			return;
		}

		if (!formEmail.trim()) {
			formError = 'Email is required';
			return;
		}

		try {
			formSubmitting = true;
			formError = null;
			
			await createGuest({
				name: formName,
				email: formEmail,
				phone: formPhone,
				status: formStatus
			});

			// Reset form
			formName = '';
			formEmail = '';
			formPhone = '';
			formStatus = 'pending';
			showForm = false;

			// Reload guests
			await loadGuests();
		} catch (e) {
			formError = e instanceof Error ? e.message : 'Failed to create guest';
		} finally {
			formSubmitting = false;
		}
	}

	async function handleDelete(id: number) {
		if (!confirm('Are you sure you want to delete this guest?')) {
			return;
		}

		try {
			await deleteGuest(id);
			await loadGuests();
		} catch (e) {
			alert(e instanceof Error ? e.message : 'Failed to delete guest');
		}
	}

	function getStatusBadgeClass(status: string): string {
		switch (status) {
			case 'attending':
				return 'bg-green-100 text-green-800';
			case 'declined':
				return 'bg-red-100 text-red-800';
			default:
				return 'bg-yellow-100 text-yellow-800';
		}
	}

	onMount(() => {
		loadGuests();
	});

	// Reload when filter changes
	$effect(() => {
		statusFilter; // Track the statusFilter variable
		loadGuests();
	});
</script>

<div class="min-h-screen bg-gray-50 py-8">
	<div class="max-w-6xl mx-auto px-4">
		<!-- Header -->
		<div class="mb-8">
			<h1 class="text-3xl font-bold text-gray-900 mb-2">Event Guest List Manager</h1>
			<p class="text-gray-600">Manage your event guests and track RSVPs</p>
		</div>

		<!-- Actions Bar -->
		<div class="bg-white rounded-lg shadow-sm p-4 mb-6">
			<div class="flex flex-wrap gap-4 items-center justify-between">
				<div class="flex items-center gap-4">
					<label class="flex items-center gap-2">
						<span class="text-sm font-medium text-gray-700">Filter by status:</span>
						<select
							bind:value={statusFilter}
							class="border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
						>
							<option value="">All Guests</option>
							<option value="pending">Pending</option>
							<option value="attending">Attending</option>
							<option value="declined">Declined</option>
						</select>
					</label>
				</div>

				<button
					onclick={() => showForm = !showForm}
					class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm font-medium transition-colors"
				>
					{showForm ? 'Cancel' : '+ Add Guest'}
				</button>
			</div>
		</div>

		<!-- Add Guest Form -->
		{#if showForm}
			<div class="bg-white rounded-lg shadow-sm p-6 mb-6">
				<h2 class="text-xl font-semibold text-gray-900 mb-4">Add New Guest</h2>
				
				<form onsubmit={handleSubmit} class="space-y-4">
					{#if formError}
						<div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-md text-sm">
							{formError}
						</div>
					{/if}

					<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
						<div>
							<label for="name" class="block text-sm font-medium text-gray-700 mb-1">
								Name *
							</label>
							<input
								type="text"
								id="name"
								bind:value={formName}
								class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
								required
							/>
						</div>

						<div>
							<label for="email" class="block text-sm font-medium text-gray-700 mb-1">
								Email *
							</label>
							<input
								type="email"
								id="email"
								bind:value={formEmail}
								class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
								required
							/>
						</div>

						<div>
							<label for="phone" class="block text-sm font-medium text-gray-700 mb-1">
								Phone
							</label>
							<input
								type="tel"
								id="phone"
								bind:value={formPhone}
								class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="status" class="block text-sm font-medium text-gray-700 mb-1">
								RSVP Status
							</label>
							<select
								id="status"
								bind:value={formStatus}
								class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
							>
								<option value="pending">Pending</option>
								<option value="attending">Attending</option>
								<option value="declined">Declined</option>
							</select>
						</div>
					</div>

					<div class="flex gap-3">
						<button
							type="submit"
							disabled={formSubmitting}
							class="bg-blue-600 hover:bg-blue-700 disabled:bg-blue-400 text-white px-6 py-2 rounded-md font-medium transition-colors"
						>
							{formSubmitting ? 'Adding...' : 'Add Guest'}
						</button>
						<button
							type="button"
							onclick={() => showForm = false}
							class="bg-gray-200 hover:bg-gray-300 text-gray-700 px-6 py-2 rounded-md font-medium transition-colors"
						>
							Cancel
						</button>
					</div>
				</form>
			</div>
		{/if}

		<!-- Guest List -->
		<div class="bg-white rounded-lg shadow-sm overflow-hidden">
			{#if loading}
				<div class="p-8 text-center text-gray-500">
					Loading guests...
				</div>
			{:else if error}
				<div class="p-8 text-center text-red-600">
					{error}
				</div>
			{:else if guests.length === 0}
				<div class="p-8 text-center text-gray-500">
					No guests found. Add your first guest to get started!
				</div>
			{:else}
				<div class="overflow-x-auto">
					<table class="w-full">
						<thead class="bg-gray-50 border-b border-gray-200">
							<tr>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Name
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Email
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Phone
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Status
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Plus Ones
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Dietary
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Notes
								</th>
								<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
									Actions
								</th>
							</tr>
						</thead>
						<tbody class="divide-y divide-gray-200">
							{#each guests as guest (guest.id)}
								<tr class="hover:bg-gray-50">
									<td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
										{guest.name}
									</td>
									<td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
										{guest.email}
									</td>
									<td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
										{guest.phone || '-'}
									</td>
									<td class="px-6 py-4 whitespace-nowrap">
										<span class={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${getStatusBadgeClass(guest.status)}`}>
											{guest.status}
										</span>
									</td>
									<td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
										{guest.plus_ones || 0}
									</td>
									<td class="px-6 py-4 text-sm text-gray-600 max-w-xs truncate" title={guest.dietary_restrictions || '-'}>
										{guest.dietary_restrictions || '-'}
									</td>
									<td class="px-6 py-4 text-sm text-gray-600 max-w-xs truncate" title={guest.notes || '-'}>
										{guest.notes || '-'}
									</td>
									<td class="px-6 py-4 whitespace-nowrap text-sm">
										<button
											onclick={() => handleDelete(guest.id)}
											class="text-red-600 hover:text-red-800 font-medium transition-colors"
										>
											Delete
										</button>
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			{/if}
		</div>

		<!-- Stats Footer -->
		{#if !loading && guests.length > 0}
			<div class="mt-6 bg-white rounded-lg shadow-sm p-4">
				<div class="flex flex-wrap gap-6 text-sm">
					<div>
						<span class="text-gray-600">Total Guests:</span>
						<span class="font-semibold text-gray-900 ml-2">{guests.length}</span>
					</div>
					<div>
						<span class="text-gray-600">Attending:</span>
						<span class="font-semibold text-green-600 ml-2">
							{guests.filter(g => g.status === 'attending').length}
						</span>
					</div>
					<div>
						<span class="text-gray-600">Pending:</span>
						<span class="font-semibold text-yellow-600 ml-2">
							{guests.filter(g => g.status === 'pending').length}
						</span>
					</div>
					<div>
						<span class="text-gray-600">Declined:</span>
						<span class="font-semibold text-red-600 ml-2">
							{guests.filter(g => g.status === 'declined').length}
						</span>
					</div>
				</div>
			</div>
		{/if}
	</div>
</div>

