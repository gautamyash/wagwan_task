# Task: Build a Public RSVP Page

## Your Mission

Design and implement a public-facing RSVP page where guests can register for an event.

## What to Build

### Public RSVP Page (`/rsvp` route)

Create a new page that allows anyone to RSVP for the event.

**Required Features:**

1. **Event Information Display**
   - Event title
   - Date and time
   - Location or description
   - Make it visually appealing and UX friendly

2. **RSVP Form**
   - Name
   - Email 
   - Phone 
   - RSVP Status selection:
     - "Yes, I'll be there!" (attending)
     - "Maybe" (pending)
     - "Sorry, can't make it" (declined)

3. **Form Behavior**
   - Validate all inputs before submission
   - Submit to the existing API
   - Show success message after submission
   - Show friendly error messages if something goes wrong
   - Clear form after successful submission

4. **Design Requirements**
   - Mobile-responsive
   - Use TailwindCSS
   - Different look from the admin interface
   - Professional and polished

## Important Notes

⚠️ **The existing codebase is functional but not perfect.**

While implementing your RSVP page, you may encounter issues with the existing code. Part of this task is to:
- Identify any bugs or problems
- Fix them as you go
- Document what you fixed and why

## Bonus Points (Optional)

If you have extra time and want to go further:

- Add duplicate email detection (prevent same email registering twice)
- Add a "View Your RSVP" feature (lookup by email)
- Add form animations or transitions
- Display attendee count ("X people attending")
- Add dietary restrictions or notes field
- Implement better error handling

## Testing Your Work

Make sure to test:
- ✅ Form validation works (required fields, email format)
- ✅ Successful submission saves to database
- ✅ Success message displays correctly
- ✅ Error messages display if API fails
- ✅ Works on mobile (responsive design)
- ✅ Admin page can see newly created guests
- ✅ All existing functionality still works

## Submission

When complete, submit:

1. **Your code** (zip file or private GitHub repo)

2. **SUBMISSION.md** with:
   - Bugs you found and fixed
   - Features you implemented
   - Challenges you faced
   - How to test your work
   - What you'd improve with more time

## Evaluation

We're looking for:
- **Working code** - Does it work as specified?
- **Code quality** - Is it clean, readable, and maintainable?
- **Problem-solving** - Did you identify and fix issues?
- **Design** - Is the RSVP page well-designed and user-friendly?
- **Testing** - Does everything work correctly?

## Time Estimate

This should take approximately 24-48 hours depending on your experience level and how many bugs you encounter.

Take your time, write clean code, and don't hesitate to ask questions if you get stuck!

Good luck! 🚀

