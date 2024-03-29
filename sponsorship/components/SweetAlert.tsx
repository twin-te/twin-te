import Swal from 'sweetalert2';
import withReactContent from 'sweetalert2-react-content';

const _MySwal = withReactContent(Swal);

// sweet alertの共通の設定を書く
export const SweetModal = _MySwal.mixin({
	confirmButtonColor: '#3085d6',
	background: 'whitesmoke'
});
