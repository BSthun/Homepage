<script lang="ts">
	import moment from 'moment'
	import { Link } from 'svelte-navigator'

	export let item: object
</script>

<div class="section-meta">
	{#if item.album != null}
		<Link to={'/photo/album/' + item.album.slug}>
			<div class="r-1 bordered">
				<i class="las la-arrow-left" />
				<div class="">
					<div class="r-1 subtitle">
						<i class="las la-images" />
						Album
					</div>
					<h4>{item.album.name}</h4>
				</div>
			</div>
		</Link>
	{/if}
	<h1>{item.section.title}</h1>
	<h4>{item.section.subtitle}</h4>
	<div class="row">
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

	.bordered {
		width: fit-content;
		line-height: 1.3;
		padding: 6px 12px;
		border-radius: 16px;
		border: 1px solid white;
		transition: background-color 0.2s ease-in-out;

		&:hover {
			background-color: rgba(255, 255, 255, 0.1);
		}
	}

	.subtitle {
		margin: 0;
		font-size: 0.7rem;
		font-weight: 100;

		i {
			font-size: 1em;
		}
	}
</style>
