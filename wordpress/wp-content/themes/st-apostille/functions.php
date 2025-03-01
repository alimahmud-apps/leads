<?php

/* 
** Sets up theme defaults and registers support for various WordPress features
*/

define('IS_ST_FREEMIUM', 'st-apostille');

function st_apostille_demo_importer_setup() {
		
	// Add default posts and comments RSS feed links to head
	add_theme_support( 'automatic-feed-links' );

	// Let WordPress manage the document title for us
	add_theme_support( 'title-tag' );

	// Enable support for Post Thumbnails on posts and pages
	add_theme_support( 'post-thumbnails' );

	// Custom Logo
	add_theme_support( 'custom-logo', [
		'height'      => 100,
		'width'       => 350,
		'flex-height' => true,
		'flex-width'  => true,
	] );

	// Custom Header
	add_theme_support( 'custom-header', array(
		'default-image'          => '',
		'default-text-color'     => '#FF7900',
		'width'                  => 1920,
		'height'                 => 1080,
		'flex-width'             => true,
		'flex-height'            => true,
		'wp-head-callback'       => 'st_apostille_header_style', // Callback for styling
	) );

	// Add theme support for Custom Background.
	add_theme_support( 'custom-background', ['default-color' => ''] );

	// Set the default content width.
	$GLOBALS['content_width'] = 960;

	// This theme uses wp_nav_menu() in one location.
	register_nav_menus(
		array(
			'primary' => esc_html__( 'Primary', 'st-apostille' ),
		)
	);

	// Switch default core markup for search form, comment form, and comments to output valid HTML5
	add_theme_support( 'html5', array(
		'comment-form',
		'comment-list',
		'gallery',
		'caption',
	) );

	// Gutenberg Embeds
	add_theme_support( 'responsive-embeds' );
	add_theme_support( 'wp-block-styles' );

	// Gutenberg Widget Images
	add_theme_support( 'align-wide' );

	// WooCommerce in general.
	add_theme_support( 'woocommerce' );

	// Zoom.
	add_theme_support( 'wc-product-gallery-zoom' );
	// Lightbox.
	add_theme_support( 'wc-product-gallery-lightbox' );
	// Swipe.
	add_theme_support( 'wc-product-gallery-slider' );

	add_editor_style( array( '/assets/css/editor-style.css' ) );
}

add_action( 'after_setup_theme', 'st_apostille_demo_importer_setup' );

/*
** Enqueue scripts and styles
*/
function st_apostille_demo_importer_scripts() {

	// Theme Stylesheet
	wp_enqueue_style( 'bootstrap-css', get_template_directory_uri() . '/assets/css/bootstrap.css', array(), '4.5.0' );

	// Remove prefix from Font Awesome
	wp_enqueue_style(
		'fontawesome-css', // Corrected handle
		get_template_directory_uri() . '/assets/css/fontawesome-all.css',
		array(),
		'4.5.0'
	);
	
	wp_enqueue_style( 'st-apostille-style', get_stylesheet_uri(), array(), '1.0' );

	// Comment reply link
	if ( is_singular() && comments_open() && get_option( 'thread_comments' ) ) {
		wp_enqueue_script( 'comment-reply' );
	}
	
	wp_enqueue_script('st-apostille-custom-script', get_template_directory_uri() . '/assets/js/script.js', array('jquery'), '1.0', true);
	
	wp_enqueue_script('st-apostille-navigation', get_template_directory_uri() . '/assets/js/navigation.js', false, '1.0', true );
}
add_action( 'wp_enqueue_scripts', 'st_apostille_demo_importer_scripts' );

function st_apostille_enqueue_admin_script() {
    // Enqueue the custom admin script
    wp_enqueue_script('st-apostille-custom-admin-script', get_template_directory_uri() . '/assets/js/custom-admin-script.js', array('jquery'), '1.0', true);
}
add_action('admin_enqueue_scripts', 'st_apostille_enqueue_admin_script');

/*
** Notices
*/
require_once get_parent_theme_file_path( '/inc/activation/class-welcome-notice.php' );
require_once get_parent_theme_file_path( '/inc/activation/class-rating-notice.php' );

add_action( 'after_switch_theme', 'st_apostille_activation_time');
add_action('after_setup_theme', 'st_apostille_activation_time');
    
function st_apostille_activation_time() {
	if ( false === get_option( 'st_apostille_activation_time' ) ) {
		add_option( 'st_apostille_activation_time', strtotime('now') );
	}
}

function st_apostille_custom_loader(){
?>
    <div id="st-apostille-loader-container">
        <div id="st-apostille-custom-loader"></div>
    </div>
<?php
}
add_action('wp_head', 'st_apostille_custom_loader');

/*
** Additions for Header Text Color and Header Image
*/

/**
 * Custom styles for Header Text Color.
 */
function st_apostille_header_style() {
    $st_apostille_header_text_color = get_header_textcolor();

    // Get the default header text color defined in the theme support array.
    $st_apostille_default_header_text_color = get_theme_support( 'custom-header', 'default-text-color' );

    // Exit if default header text color is being used.
    if ( $st_apostille_header_text_color === $st_apostille_default_header_text_color ) {
        return;
    }

    // Custom CSS for header text color.
    ?>
    <style type="text/css">
        <?php if ( ! display_header_text() ) : ?>
            .site-title,
            .site-description {
                display: none;
            }
        <?php else : ?>
            .site-title a,
            .site-description {
                color: #<?php echo esc_attr( $st_apostille_header_text_color ); ?>;
            }
        <?php endif; ?>
    </style>
    <?php
}

/**
 * Customize live preview for Header Text Color.
 */
function st_apostille_customize_live_preview() {
    wp_enqueue_script(
        'st-apostille-customizer',
        get_template_directory_uri() . '/assets/js/customizer.js',
        array( 'customize-preview' ),
        '1.0',
        true
    );
}
add_action( 'customize_preview_init', 'st_apostille_customize_live_preview' );
