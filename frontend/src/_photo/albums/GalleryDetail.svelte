<script lang="ts">
	import moment from 'moment'
	import { trackLog } from '../../utils/api/track'
	import { getContext, onDestroy } from 'svelte'
	import Zoom from 'svelte-zoom'
	import CircularProgress from '../../components/indicator/CircularProgress.svelte'

	const gallery = getContext('gallery')
	let loading = true
	let plane = false
	let image = new Image()
	image.onload = () => {
		loading = false
	}

	const forward = (i) => {
		if (i < 0 && $gallery.expand != 0) {
			$gallery.expand -= 1
		}
		if (i > 0 && $gallery.expand != $gallery.items.length - 1) {
			$gallery.expand += 1
		}
		if (i > 0 && $gallery.expand == $gallery.items.length - 1 && $gallery.expand < $gallery.count) {
			$gallery.fetch()
		}
		trackLog('gallery/expand', null, $gallery.items[$gallery.expand].id)
	}

	const download = (id, url, mode) => {
		trackLog('gallery/download/' + mode, null, id)
		window.location.href = url + '?download=1'
	}

	let isClosing = false
	const onHashChange = () => {
		if (!window.location.hash) {
			$gallery.expand = null
			if (isClosing) {
				isClosing = false
			} else {
				trackLog('gallery/close/hash', null, null)
			}
		}
	}

	onDestroy(() => {
		window.removeEventListener('hashchange', onHashChange)
	})

	$: item = $gallery.items[$gallery.expand]
	$: {
		window.addEventListener('hashchange', onHashChange)
	}
	$: {
		if (item !== undefined) {
			loading = true
			plane = false
			image.src = item && item.root + item.image_path
		}
	}
	$: {
		if ($gallery.expand != null) {
			$gallery.loading = true
			document.querySelector('body').style.overflow = 'hidden'
			window.location.hash = 'plane'
		} else {
			document.querySelector('body').style.overflow = 'scroll'
		}
	}
</script>

{#if $gallery.expand != null}
	<div class="gallery-detail">
		<div class="img-wrapper">
			{#if loading === true}
				<div class="center">
					<CircularProgress />
				</div>
			{:else}
				<Zoom alt={item.title} src={item.root + item.image_path} />
			{/if}
			<span class="material-symbols-outlined back" on:click={() => forward(-1)}>arrow_back</span>
			<span class="material-symbols-outlined forward" on:click={() => forward(1)}>arrow_forward</span>
		</div>
		<div class={plane ? 'plane plane-toggled' : 'plane'}>
			<div class="plane-close">
				{#if !plane}
					<div
						class="plane-detail-toggle"
						on:click={() => {
							plane = !plane
							trackLog('gallery/plane', null, plane)
						}}
					>
						<p>VIEW DETAIL / DOWNLOAD</p>
					</div>
				{:else}
					<div />
				{/if}
				<span
					class="material-symbols-outlined"
					on:click={() => {
						if (plane) {
							plane = !plane
							trackLog('gallery/plane', null, plane)
							return
						}
						trackLog('gallery/close', null, null)
						isClosing = true
						history.back()
					}}
				>
					{plane ? 'keyboard_arrow_down' : 'close'}
				</span>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">schedule</span>
				<div>
					<p>Timestamp</p>
					<p>{moment(item.exif.t).format('MMM D YYYY, h:mm a')}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">camera</span>
				<div>
					<p>Aperture</p>
					<p>ƒ/{item.exif.apt}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">shutter_speed</span>
				<div>
					<p>Shutter Speed</p>
					<p>{item.exif.ss}s</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">crop</span>
				<div>
					<p>Focal Length</p>
					<p>{item.exif.fl}mm ({item.exif.fl_ff}mm FF)</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">deblur</span>
				<div>
					<p>ISO Speed</p>
					<p>{item.exif.iso}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">photo_camera</span>
				<div>
					<p>Camera</p>
					<p>{item.exif.c}</p>
				</div>
			</div>
			<div class="plane-info">
				<span class="material-symbols-outlined">lens</span>
				<div>
					<p>Lens</p>
					<p>{item.exif.l}</p>
				</div>
			</div>
			<div class="plane-download" on:click={() => download(item.id, item.root + item.image_path, 'img')}>
				<span class="material-symbols-outlined">image</span>
				<p>Download Image</p>
			</div>
			<div class="plane-download" on:click={() => download(item.id, item.root + item.raw_path, 'raw')}>
				<span class="material-symbols-outlined">add_photo_alternate</span>
				<p>Download Raw</p>
			</div>
		</div>
	</div>
{/if}

<style lang="scss">
	@import 'src/styles/_index.scss';

	.gallery-detail {
		position: fixed;
		inset: 0;
		z-index: 10000;
		background-color: #1e2118ea;
		display: flex;

		@include breakpoint('md', 'dn') {
			flex-direction: column;
		}
	}

	.img-wrapper {
		position: relative;
		flex: 1;
		overflow: hidden;

		@include breakpoint('md', 'up') {
			border-radius: 8px;
			margin: 12px 0 12px 12px;
		}

		@include breakpoint('md', 'dn') {
			margin-bottom: 48px;
		}

		.material-symbols-outlined {
			top: calc(50% - 12px);
			position: absolute;
			z-index: 200;
			color: $color-grey-800;
			background-color: $color-grey-100;
			@include tweak-clickable-icon(rgba($color-grey-200, 0.7));
		}

		.back {
			left: 4px;
		}

		.forward {
			right: 4px;
		}

		:global(img) {
			-webkit-touch-callout: none;
		}
	}

	.plane {
		display: flex;
		flex-direction: column;
		background-color: white;
		color: $color-grey-900;
		transition: top 300ms ease-in-out;

		@include breakpoint('md', 'up') {
			width: 300px;
			border-radius: 8px;
			margin: 12px;
		}

		@include breakpoint('md', 'dn') {
			position: absolute;
			z-index: 201;
			top: calc(100% - 48px);
			width: 100%;
			height: 100%;
		}
	}

	.plane-toggled {
		top: 0;
	}

	.plane-close {
		padding: 8px;
		display: flex;
		justify-content: space-between;
		align-items: center;

		> span {
			@include tweak-clickable-icon($color-grey-300);
		}
	}

	.plane-detail-toggle {
		border-radius: 4px;
		border: #546e7a 1px solid;
		padding: 4px 8px;

		p {
			font-size: 12px;
		}

		@include breakpoint('md', 'up') {
			visibility: hidden;
		}
	}

	.plane-info {
		display: flex;
		align-items: center;
		padding: 12px;
		gap: 12px;

		> div {
			flex: 1;

			p:first-child {
				font-size: 12px;
				color: $color-bluegrey-700;
			}

			p:last-child {
				font-size: 16px;
				color: $color-grey-900;
			}
		}
	}

	.plane-download {
		margin: 12px 12px 0 12px;
		padding: 6px 12px;
		border: #546e7a 1px solid;
		border-radius: 4px;
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 6px;
		font-size: 16px;
		@include tweak-clickable($color-default-clickable);
	}

	.center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
	}
</style>
