export default setup() {
    document.body.addEventListener('showRoster', handleShowEvent)

    /* Shows the roster after the first player is added.
     * @returns { void } 
     */
    function handleShowEvent() {
        var node = document.querySelector('.roster-table');
        if (node.classList.contains('hidden')) {
            node.classList.remove('hidden');
        }
    };
};
