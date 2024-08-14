export default (function () {
    document.body.addEventListener('showRoster', handleShowEvent)
    document.body.addEventListener('downloadCsv', handleDownloadFile)

    /* Shows the roster after the first player is added.
     * @returns { void }
     */
    function handleShowEvent() {
        var node = document.querySelector('.roster-table');
        if (node.classList.contains('hidden')) {
            node.classList.remove('hidden');
        }
    };

    /* Open a new tab/window to download the generated CSV file
     * @returns { void }
     */
    function handleDownloadFile(event) {
        window.open(`/download/${event.detail.value}`);
    }
})();
