<script lang="ts">
	import moment from 'moment'

	export let item: object
</script>

<div class="section-meta">
	<h1>{item.section.title}</h1>
	<h4>{item.section.subtitle}</h4>
	<div class="row">
		{#if item.album != null}
			<div class="r-1">
				<i class="las la-images" />
				<h4>{item.album.name}</h4>
			</div>
		{/if}
		<div class="r-1">
			<i class="las la-image" />
			<h4>{item.section.photo_count}</h4>
		</div>
		<div class="r-1">
			<i class="las la-clock" />
			<h4>{moment(item.section.date).calendar()}</h4>
		</div>
	</div>
</div>

<style lang="scss">
	@import 'src/styles/_index.scss';

	.section-meta {
		display: flex;
		flex-direction: column;
		gap: 4px;
	}

	.row {
		display: flex;

		@include breakpoint('md', 'dn') {
			flex-direction: column;
		}
	}
</style>
