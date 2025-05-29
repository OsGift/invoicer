function addItem() {
    const itemContainer = document.getElementById('items');
    const itemEntry = document.createElement('div');
    itemEntry.className = 'item-entry';
    itemEntry.innerHTML = `
        <input type="text" name="itemDescription[]" placeholder="Item Description" required>
        <input type="number" name="itemQuantity[]" placeholder="Quantity" required>
        <input type="number" step="0.01" name="itemPrice[]" placeholder="Price" required>
    `;
    itemContainer.appendChild(itemEntry);
}
