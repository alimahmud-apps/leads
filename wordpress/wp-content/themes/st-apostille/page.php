<?php get_header(); ?>

<div id="skip-link-target" class="mt-5">

	<!-- Main Container -->
	<div class="main-container container page-single">
		
		<article id="page-<?php the_ID(); ?>" <?php post_class(); ?>>

			<?php
			if ( have_posts() ) :
			while ( have_posts() ) : the_post();

				if ( has_post_thumbnail() ) {
					echo '<div class="post-media">';
						the_post_thumbnail();
					echo '</div>';
				}

				if ( get_the_title() !== '' ) {
					echo '<header class="post-header">';
						echo '<h1 class="page-title">'. get_the_title() .'</h1>';
					echo '</header>';
				}

				echo '<div class="post-content">';
					the_content('');

					$st_apostille_defaults = array(
						'before' => '<p class="single-pagination">'. esc_html__( 'Pages:', 'st-apostille' ),
						'after' => '</p>'
					);

					wp_link_pages( $st_apostille_defaults );
				echo '</div>';

			endwhile;
			endif;
			?>

		</article>

		<?php 
		if ( comments_open() || get_comments_number() ) {
			echo '<div class="comments-area" id="comments">';
				comments_template( '', true );
			echo '</div>';
		}
		?>

	</div><!-- .main-container -->
</div>

<?php get_footer(); ?>