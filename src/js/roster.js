document.body.addEventListener('showRoster', handleShowEvent)

/* Shows the roster after the first player is added.
 * @returns { void } 
 */
function handleShowEvent() {
    document.querySelector('.roster-container').classList.Remove('hidden');
};
