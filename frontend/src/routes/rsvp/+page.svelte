<script lang="ts">
    import { onMount } from "svelte";

    let event = {
        title: "Wagwan Launch Party 🎉",
        description:
            "Join us for an evening of fun, networking, and celebration!",
        date: "December 15, 2025",
        time: "6:00 PM - 10:00 PM",
        location: "Wagwan HQ, Pune, India",
    };

    let name = "";
    let email = "";
    let phone = "";
    let status = "attending";
    let notes = "";
    let plus_ones = 0;
    let dietary_restrictions = "";

    let isSubmitting = false;
    let successMessage = "";
    let errorMessage = "";

    const apiURL = "http://localhost:8080/api/guests";

    const validateForm = () => {
        if (!name.trim() || !email.trim() || !phone.trim()) {
            errorMessage = "Please fill in all required fields.";
            return false;
        }
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            errorMessage = "Please enter a valid email address.";
            return false;
        }
        return true;
    };

    const submitRSVP = async () => {
        if (!validateForm()) return;

        isSubmitting = true;
        errorMessage = "";
        successMessage = "";

        try {
            const res = await fetch(apiURL, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    name,
                    email,
                    phone,
                    status,
                    notes,
                    plus_ones: Number(plus_ones),
                    dietary_restrictions,
                    event_id: 1,
                }),
            });

            if (!res.ok) throw new Error("Submission failed");

            successMessage = "🎉 Thank you! Your RSVP has been recorded.";
            name = email = phone = notes = dietary_restrictions = "";
            plus_ones = 0;
            status = "attending";
        } catch (err) {
            errorMessage = "Something went wrong. Please try again.";
        } finally {
            isSubmitting = false;
        }
    };
</script>

<!-- PAGE -->
<div class="min-h-screen flex flex-col items-center justify-center bg-gray-50 py-10 px-6">
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

        <!-- RSVP FORM -->
        <form on:submit|preventDefault={submitRSVP} class="flex flex-col gap-4">
            <input type="text" bind:value={name} placeholder="Full Name" class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none" required />
            <input type="email" bind:value={email} placeholder="Email Address" class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none" required />
            <input type="tel" bind:value={phone} placeholder="Phone Number" class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none" required />

            <select bind:value={status} class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none">
                <option value="attending">✅ Yes, I’ll be there!</option>
                <option value="pending">🤔 Maybe</option>
                <option value="declined">❌ Sorry, can’t make it</option>
            </select>

            <input type="number" min="0" bind:value={plus_ones} placeholder="Number of plus ones" class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none" />

            <textarea bind:value={dietary_restrictions} placeholder="Dietary restrictions (optional)" class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"></textarea>

            <textarea bind:value={notes} placeholder="Additional notes (optional)" class="border rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-500 outline-none"></textarea>

            <button type="submit" class="bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-lg py-2 transition disabled:opacity-50" disabled={isSubmitting}>
                {isSubmitting ? "Submitting..." : "Submit RSVP"}
            </button>

            {#if successMessage}
                <p class="text-green-600 text-center font-medium mt-2">{successMessage}</p>
            {/if}
            {#if errorMessage}
                <p class="text-red-600 text-center font-medium mt-2">{errorMessage}</p>
            {/if}
        </form>
    </div>
</div>
