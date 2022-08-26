<script lang="ts">
	let props: String;
	$: ({ ...props } = $$props);
</script>

<div {...props} class="container">
	<slot></slot>
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';
	
	.container {
		width: calc(100% - 32px);
		height: 100%;
		max-width: map-get($breakpoints, 'xl');
		margin: 0 auto;
		padding: 16px;
	}
</style>
