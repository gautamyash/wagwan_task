<script lang="ts">
    export let form:
        | {
              success?: boolean;
              message?: string;
              error?: string;
              values?: Record<string, unknown>;
          }
        | undefined;

    const event = {
        title: "Wagwan Launch Party 🎉",
        description:
            "Join us for an evening of fun, networking, and celebration!",
        date: "December 15, 2025",
        time: "6:00 PM - 10:00 PM",
        location: "Wagwan HQ, Pune, India",
    };
</script>

<!-- PAGE -->
<div
    class="min-h-screen flex flex-col items-center justify-center bg-gray-50 py-10 px-6"
>
    <div class="w-full max-w-lg bg-white rounded-2xl shadow-lg p-8">
        <!-- Event Info -->
        <div class="text-center mb-6">
            <h1 class="text-3xl font-bold text-gray-900">{event.title}</h1>
            <p class="text-gray-600 mt-2">{event.description}</p>
            <p class="mt-3 font-medium text-gray-800">
                📅 {event.date} | 🕕 {event.time}
            </p>
            <p class="text-gray-600">📍 {event.location}</p>
        </div>

        <div class="border-t my-6"></div>

        <!-- RSVP FORM (SvelteKit server action) -->
        <form method="post" class="flex flex-col gap-4">
            <input
                name="name"
                type="text"
                placeholder="Full Name"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
                required
            />
            <input
                name="email"
                type="email"
                placeholder="Email Address"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
                required
            />
            <input
                name="phone"
                type="tel"
                placeholder="Phone Number"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
                required
            />

            <select
                name="status"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
            >
                <option value="attending">✅ Yes, I’ll be there!</option>
                <option value="pending">🤔 Maybe</option>
                <option value="declined">❌ Sorry, can’t make it</option>
            </select>

            <input
                name="plus_ones"
                type="number"
                min="0"
                placeholder="Number of plus ones"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
            />

            <textarea
                name="dietary_restrictions"
                placeholder="Dietary restrictions (optional)"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
            ></textarea>

            <textarea
                name="notes"
                placeholder="Additional notes (optional)"
                class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"
            ></textarea>

            <button
                type="submit"
                class="bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-lg py-2 transition"
            >
                Submit RSVP
            </button>

            {#if form?.message}
                <p class="text-green-600 text-center font-medium mt-2">
                    {form.message}
                </p>
            {/if}
            {#if form?.error}
                <p class="text-red-600 text-center font-medium mt-2">
                    {form.error}
                </p>
            {/if}
        </form>
    </div>
</div>
