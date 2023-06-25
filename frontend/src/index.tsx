import Router from 'preact-router';
import { h, render } from 'preact';
import { useEffect, useReducer } from 'preact/hooks';

const Main = () => (
  <Router>
    <Game path="/game/:game" />
    <Card path="/card/:card" />
  </Router>
);

interface GameProps {
    path: string,
    game?: string,
}

interface GameState {
    title: string,
    squares: any[],
    loading: boolean,
}

const gameReducer = (state: GameState, action: GameState): GameState => {
    return {
        ...state,
        ...action,
    }
}

const Game = (props: GameProps) => {
    const [game, dispatch] = useReducer(gameReducer, {
        title: '',
        loading: true,
        squares: [],
    });

    useEffect(() => {
        fetch(`/api/game/${props.game}`).then((res) => {
            return res.json();
        }).then((json) => {
            dispatch({
                title: json.Title,
                squares: json.Squares,
                loading: false,
            })
        });
    }, []);

    const message = game.loading ? 'Loading' : game.title;

    return <div>
        <h1>{message}</h1>
        <ul className="game">
            {game.squares.map((s: any) => {
                let className = s.Selected ? 'selected' : '';
                className = s.Needed ? className + ' needed' : className;
                return <li className={className}>{s.Title}</li>;
            })}
        </ul>
    </div>;
}

interface CardState {
    title: string,
    squares: any[],
    loading: boolean,
}

const cardReducer = (state: CardState, action: CardState): CardState => {
    return {
        ...state,
        ...action,
    }
}

interface CardProps {
    path: string,
    card?: string,
}

const Card = (props: CardProps) => {
    const [card, dispatch] = useReducer(cardReducer, {
        title: '',
        loading: true,
        squares: [],
    });

    useEffect(() => {
        fetch(`/api/card/${props.card}`).then((res) => {
            return res.json();
        }).then((json) => {
            console.log(json);
            dispatch({
                title: 'Bingo Card',
                squares: json.Squares,
                loading: false,
            })
        });
    }, []);

    const message = card.loading ? 'Loading' : card.title;

    return <div>
        <h1>{message}</h1>
        <ul className="card">
            {card.squares.map((s: any) => {
                const className = s.Selected ? 'selected' : '';
                return <li className={className}>{s.Title}</li>;
            })}
        </ul>
    </div>;
}

render(<Main />, document.body);
