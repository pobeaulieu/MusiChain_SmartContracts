import load from '../images/loading.svg';

const Button = (props : any) => {
  return (
    <button className="submit-btn"  >
      {!props.loading ? props.text : <img src={load} alt="load" />}
    </button>
  )
}

export default Button