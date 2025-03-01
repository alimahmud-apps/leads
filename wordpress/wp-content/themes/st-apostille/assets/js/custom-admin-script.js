jQuery(document).ready(function ($) {
  // Find the elements with the class 'align-items-start' and append HTML content
  $('.parent-import-container').each(function () {
    // Create HTML content
    var newContent = `
    <div class='st-apostille-compare-feature'>
        <div class='st-apostille-portal-btn'>
            <a href="https://striviothemes.com/product/premium-apostille-services-wordpress-theme/" class="button button-primary st-apostille-buy-now" target="_blank">Buy Now</a>
            <a href="https://striviothemes.com/demo/st-apostille-pro" class="button button-primary st-apostille-view-demo" target="_blank">Demo</a>
        </div>
        <h2 class='st-apostille-feature-comparison'>Feature Comparison</h2>
        <table class="st-apostille-compare-table">
            <tr>
                <th class="st-apostille-compare-th">Features</th>
                <th class="st-apostille-compare-th-pro">Pro</th>
                <th class="st-apostille-compare-th">Free</th>
            </tr>
            <tr>
                <td class="st-compare-td">Animations</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">Services Custom Posttype</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">Booking Form</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">About Page</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">Contact Page</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">Services Page</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">Blog Page</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">More Sections</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            <tr>
                <td class="st-compare-td">Before After Image</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
            
            <tr>
                <td class="st-compare-td">Detailed Documentation</td>
                <td class="st-apostille-compare-td-pro"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/available.png" alt="Available"></td>
                <td class="st-apostille-compare-td-free"><img src="https://striviothemes.com/demo/st-demo-importer-imgs/not-available.png" alt="Not Available"></td>
            </tr>
        </table>
    </div>`
    // Append the HTML content to the element
    $(this).append(newContent);
  });
});
